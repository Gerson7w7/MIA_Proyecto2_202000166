package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func AnalisisRmdisk(comando string) {
	var lineacomand [100]string
	newcomando := strings.Split(comando, "") //aqui creo un arreglo del string asi puedo recorrer carcater por caracter
	lineacomando := ""                       //Este nos sive igual que en c++ cuando usamabamos el char y madabamos la informacion a y comparamos lo que viene

	//copy
	copy(lineacomand[:], newcomando[:]) //aqui ya lo copiamos a lineacomand para tener el arreglo

	//banderas
	flag_path := false
	//valor
	valor_path := ""
	contador := 0

	//simula un while
	for strings.Compare(lineacomand[contador], "\n") != 0 && strings.Compare(lineacomand[contador], "#") != 0 {
		//validacion de caracteres para interrupcion

		if strings.Compare(lineacomand[contador], " ") == 0 {
			contador++
			lineacomando = ""
		} else {
			lineacomando += strings.ToLower(lineacomand[contador])
			contador++
		}

		//validacion de comandos y sus valores
		if strings.Compare(lineacomando, "rmdisk") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
		} else if strings.Compare(lineacomando, ">path=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n"
			lineacomando = ""
			flag_path = true

			//simula in while
			if strings.Compare(lineacomand[contador], "\"") == 0 {

				contador++
				//simula otro while
				for strings.Compare(lineacomand[contador], "\n") != 0 {
					if strings.Compare(lineacomand[contador], "\"") == 0 {
						break
					} else {
						valor_path += lineacomand[contador]
						contador++
					}
				}
			} else {
				//simula un while
				for strings.Compare(lineacomand[contador], "\n") != 0 {
					if strings.Compare(lineacomand[contador], " ") == 0 || strings.Compare(lineacomand[contador], "\n") == 0 {
						break
					} else {
						valor_path += lineacomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_path)
			cadenaf += "Valor: " + valor_path + " \n"
		}
	}

	//proceso para eliminar disk
	flag_available := validacionDisk(valor_path) // para verificar si existe el disco a eliminar
	if flag_path == true {
		if flag_available == true {
			fmt.Println("Se eliminará este disco " + valor_path)
			cadenaf += "Se eliminará este disco " + valor_path + " \n"
			deleteDisk(valor_path)
			fmt.Println("¡Disco eliminado correctamente!")
			cadenaf += "¡Disco eliminado correctamente! \n"
		} else {
			fmt.Println("Error -> El disco no existe")
			cadenaf += "Error -> El disco no existe \n"
		}
	} else {
		fmt.Println("Error -> No está el path ")
		cadenaf += "Error -> No está el path \n"
	}
}

//Metodo de eliminacion
func deleteDisk(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal("Error al eliminar disco", err)
	}
}

//metodo para la validacion del path para eliminacion de Disco
func validacionDisk(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}