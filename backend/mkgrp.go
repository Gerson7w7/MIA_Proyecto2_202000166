package main

import (
	"fmt"
	"strings"
)

func AnalisiMkgrp(comando string) {
	fmt.Println("Entramos al mkgrp")
	var linecomand [200]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""
	// se realiza una copia del array para manejo
	copy(linecomand[:], newcomando[:])

	//banderas
	flag_name := false

	//valores
	valor_name := ""

	contador := 0
	//simula un while
	for strings.Compare(linecomand[contador], "\n") != 0 && strings.Compare(linecomand[contador], "#") != 0 {
		//validacion de caracteres para interrupcion
		if strings.Compare(linecomand[contador], " ") == 0 {
			contador++
			lineacomando = ""
			//este if sirve para validar comandos vacios y de no traer nada no validara nada de esta parte
		} else {
			lineacomando += strings.ToLower(linecomand[contador])
			contador++
			//este solo sivere para confirmar que encontro el comando mkdisk nada mas
		}

		//validacion de valores de comando
		if strings.Compare(lineacomando, "mkgrp") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
			//aqui limpio la linea comando y el contador sige en aumento para validar los demas caracteres.
		} else if strings.Compare(lineacomando, ">name=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n" //Esto es lo que se mostrara en el Front
			lineacomando = ""                              //se limpia linea
			flag_name = true                               //se activa bandera size	 ya que es comando obligatorio
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], "\"") == 0 { // si viene con comilla doble
					contador++
					//simula un while
					for strings.Compare(linecomand[contador], "\n") != 0 {
						if strings.Compare(linecomand[contador], "\"") == 0 { //finaliza path
							contador++
							break
						} else {
							valor_name += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_name += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_name)
			cadenaf += "Valor: " + valor_name + " \n"
		}
	}
	//---------------PROCESO CREACION DE GRUPOS-------------------------
	if (flag_name) {
		fmt.Println("Inicio de proceso")
		for e := Groups.Front(); e != nil; e = e.Next() {
			name := e.Value
			if valor_name == name {
				fmt.Println("Error-> El nombre de grupo ya existe, elija otro diferente: " + valor_name)
				cadenaf += "Error-> El nombre de grupo ya existe, elija otro diferente: " + valor_name + "\n"
				return;
			}
		}
		Groups.PushFront(valor_name)
		fmt.Println("Grupo agregado con éxito: " + valor_name)
		cadenaf += "Grupo agregado con éxito: " + valor_name + " \n"
		for e := Groups.Front(); e != nil; e = e.Next() {
			fmt.Println(e)
		}
	} else {
		fmt.Println("Error-> Se necesita nombre para poder crear un nuevo grupo")
		cadenaf += "Error-> Se necesita nombre para poder crear un nuevo grupo \n"
	}
}

func AnalisiRmgrp(comando string) {
	fmt.Println("Entramos al rmgrp")
	var linecomand [200]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""
	// se realiza una copia del array para manejo
	copy(linecomand[:], newcomando[:])

	//banderas
	flag_name := false

	//valores
	valor_name := ""

	contador := 0
	//simula un while
	for strings.Compare(linecomand[contador], "\n") != 0 && strings.Compare(linecomand[contador], "#") != 0 {
		//validacion de caracteres para interrupcion
		if strings.Compare(linecomand[contador], " ") == 0 {
			contador++
			lineacomando = ""
			//este if sirve para validar comandos vacios y de no traer nada no validara nada de esta parte
		} else {
			lineacomando += strings.ToLower(linecomand[contador])
			contador++
			//este solo sivere para confirmar que encontro el comando mkdisk nada mas
		}

		//validacion de valores de comando
		if strings.Compare(lineacomando, "rmgrp") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
			//aqui limpio la linea comando y el contador sige en aumento para validar los demas caracteres.
		} else if strings.Compare(lineacomando, ">name=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n" //Esto es lo que se mostrara en el Front
			lineacomando = ""                              //se limpia linea
			flag_name = true                               //se activa bandera size	 ya que es comando obligatorio
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], "\"") == 0 { // si viene con comilla doble
					contador++
					//simula un while
					for strings.Compare(linecomand[contador], "\n") != 0 {
						if strings.Compare(linecomand[contador], "\"") == 0 { //finaliza path
							contador++
							break
						} else {
							valor_name += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_name += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_name)
			cadenaf += "Valor: " + valor_name + " \n"
		} 
	}
	//---------------PROCESO ELIMINACION DE GRUPOS-------------------------
	if (flag_name) {
		fmt.Println("Inicio de proceso")
		for e := Groups.Front(); e != nil; e = e.Next() {
			name := e.Value
			if (valor_name == name) {
				Groups.Remove(e)
				fmt.Println("Grupo eliminado con éxito: " + valor_name)
				cadenaf += "Grupo eliminado con éxito: " + valor_name + " \n"
				for e := Groups.Front(); e != nil; e = e.Next() {
					fmt.Println(e)
				}
				return;
			}
		}
		fmt.Println("Grupo no encontrado, verifique el nombre: " + valor_name)
		cadenaf += "Grupo no encontrado, verifique el nombre: " + valor_name + " \n"
	} else {
		fmt.Println("Error-> Se necesita el nombre de grupo para poder eliminar un grupo")
		cadenaf += "Error-> Se necesita el nombre de grupo para poder eliminar un grupo \n"
	}
}