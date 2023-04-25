package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"log"
	
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
type User struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	Grp  string `json:"grp"`
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
	fmt.Println("Informacion: ", nweT.Usuario)
	fmt.Println("Informacion: ", nweT.Clave)
	fmt.Println("Informacion: ", nweT.Id)
	var respuesta User
	for e := Users.Front(); e != nil; e = e.Next() {
		user := e.Value.(User)
		if nweT.Usuario == user.Name && nweT.Clave == user.Pwd {
			respuesta.Name = user.Name 
			respuesta.Pwd = user.Pwd
			respuesta.Grp = ""
		}
	}
	// devolvemos la info al front
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta) //solicito el dato agregado
}

func main() {
	Users.PushFront(User{ Name: "root", Pwd: "123", Grp: "root" }) 
	Groups.PushFront("root")

	// creando el servidor y corriéndolo
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// endpoints a utilizar
	router.HandleFunc("/login",valiLogin)
	router.HandleFunc("/comandos",AnalisisCadena)
	router.HandleFunc("/reportes",analissisRepb64)

	println("Server are running on port: 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(router)))
}