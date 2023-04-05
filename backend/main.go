package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type ComandoI struct {
	Exp string `json:"exp"`
}

type ContData struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
	Id      string `json:"id"`
}

type DatoRe struct {
	Validate  bool   `json:"validate"`
	Contenido string `json:"contenido"`
}
type DatoRepIMG struct {
	Validate  bool   `json:"validate"`
	Contenido string `json:"contenido"`
	Datob64   string `json:"datob64"`
}

func AnalisisContenido(contenido string) {
	lineacomando := "" // donde se guarda el primer comando
	comandosep := strings.Split(contenido, "") 
	fmt.Println("El tamaño total del arreglo creado", len(comandosep))

	for i := 0; i < len(comandosep); i++ {
		if strings.Compare(comandosep[i], "\n") == 0 {
			fmt.Println("Se genera el comando para el analisis: ", lineacomando)
			lineacomando += (comandosep[i])
			AnalizadorComando(lineacomando)
			lineacomando = "" //receteo la linea del comando
		} else if i == (len(comandosep) - 1) {
			lineacomando += (comandosep[i])
			fmt.Println("Se genera el comando para el analisis: ", lineacomando)
			AnalizadorComando(lineacomando)

		} else {
			lineacomando += (comandosep[i]) ///le cambie el tolower
		}
	}
}

func AnalisisCadena(w http.ResponseWriter, r *http.Request) {
	// obtenemos info del front
	var nweT ComandoI
	json.NewDecoder(r.Body).Decode(&nweT)

	var respuesta DatoRe
	fmt.Println("Informacion: ", nweT.Exp)
	AnalisisContenido(nweT.Exp)
	//Aqui mando una respuesta al front
	respuesta.Validate = true
	respuesta.Contenido = cadenaf
	cadenaf = ""

	// devolvemos la info al front
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(respuesta)               
}

//para retornar la imagen base64:
func analissisRepb64(w http.ResponseWriter, r *http.Request) {
	// obtenemos info del front
	var nweT ComandoI
	json.NewDecoder(r.Body).Decode(&nweT)
	var respuesta DatoRepIMG

	fmt.Println("Informacion: ", nweT.Exp)
	AnalisisContenido(nweT.Exp)
	//Aqui mando una respuesta al front
	respuesta.Validate = repVali
	respuesta.Contenido = cadenaf
	respuesta.Datob64 = imagenFinalRep
	cadenaf = ""
	imagenFinalRep = "" //limpio la variable
	repVali = false     //aqui lo regreso a false

	// devolvemos la info al front
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(respuesta)               
}

func valiLogin(w http.ResponseWriter, r *http.Request) {
	// obtenemos info del front
	var nweT ContData
	json.NewDecoder(r.Body).Decode(&nweT)
	
	var respuesta DatoRe
	fmt.Println("Informacion: ", nweT.Usuario)
	fmt.Println("Informacion: ", nweT.Clave)
	fmt.Println("Informacion: ", nweT.Id)
	respuesta.Validate = valMontado(nweT.Id)
	
	// devolvemos la info al front
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta) //solicito el dato agregado
}

func main() {
	// creando el servidor y corriéndolo
	mux := http.NewServeMux()

	// endpoints a utilizar
	mux.HandleFunc("/login",valiLogin)
	mux.HandleFunc("/comandos",AnalisisCadena)
	mux.HandleFunc("/reportes",analissisRepb64)

	handler := cors.Default().Handler(mux)
	println("Server are running on port: 5000")
	http.ListenAndServe(":5000", handler)
}