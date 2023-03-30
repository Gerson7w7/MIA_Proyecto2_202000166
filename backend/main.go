package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type ComandoI struct {
	Exp string `json:exp`
}

type ContData struct {
	Usuario string `json:usuario`
	Clave   string `json:clave`
	Id      string `json:id`
}

type DatoRe struct {
	Validate  bool   `json:validate`
	Contenido string `json:contenido`
}
type DatoRepIMG struct {
	Validate  bool   `json:validate`
	Contenido string `json:contenido`
	Datob64   string `json:datob64`
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
	var nweT ComandoI
	var respuesta DatoRe
	reqBoy, err := ioutil.ReadAll(r.Body)
	if err != nil { //si hay errores
		fmt.Fprintf(w, "inserte datos invalidos")
	}
	json.Unmarshal(reqBoy, &nweT)

	fmt.Println("Informacion: ", nweT.Exp)
	AnalisisContenido(nweT.Exp)
	//Aqui mando una respuesta al front
	respuesta.Validate = true
	respuesta.Contenido = cadenaf
	cadenaf = ""

	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(respuesta)               
}

//para retornar la imagen base64:
func analissisRepb64(w http.ResponseWriter, r *http.Request) {
	var nweT ComandoI
	var respuesta DatoRepIMG
	reqBoy, err := ioutil.ReadAll(r.Body)
	if err != nil { //si hay errores
		fmt.Fprintf(w, "inserte datos invalidos")

	}
	json.Unmarshal(reqBoy, &nweT)

	fmt.Println("Iformacion: ", nweT.Exp)
	AnalisisContenido(nweT.Exp)
	//Aqui mando una respuesta al front
	respuesta.Validate = repVali
	respuesta.Contenido = cadenaf
	respuesta.Datob64 = imagenFinalRep
	cadenaf = ""
	imagenFinalRep = "" //limpio la variable
	repVali = false     //aqui lo regreso a false

	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(respuesta)               
}

func valiLogin(w http.ResponseWriter, r *http.Request) {
	var nweT ContData
	var respuesta DatoRe
	reqBoy, err := ioutil.ReadAll(r.Body)
	if err != nil { //si hay errores
		fmt.Fprintf(w, "inserte datos invalidos")

	}
	json.Unmarshal(reqBoy, &nweT)

	fmt.Println("Iformacion: ", nweT.Usuario)
	fmt.Println("Iformacion: ", nweT.Clave)
	fmt.Println("Iformacion: ", nweT.Id)
	respuesta.Validate = valMontado(nweT.Id)
	
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