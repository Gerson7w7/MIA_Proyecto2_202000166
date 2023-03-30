package main

import (
	"fmt"
	"strings"
)

func analisisMkfsv2(comando string) {
	fmt.Println("Realizando analisis de comando")
	var linecomand [100]string
	nuevocomando := strings.Split(comando, "")
	comandoingresado := ""
	valor_type := "full"
	valor_id := ""

	//banderas
	flag_type := false //opcional
	flag_id := false   //obligatorio

	flag_format := false
	copy(linecomand[:], nuevocomando[:])

	contador := 0 //contador para iterar en el while(for)
	for strings.Compare(linecomand[contador], "\n") != 0 && strings.Compare(linecomand[contador], "#") != 0 {
		//validacion de caractesre para interrupcion
		if strings.Compare(linecomand[contador], " ") == 0 {
			contador++
			comandoingresado = ""
		} else {
			comandoingresado += strings.ToLower(linecomand[contador])
			contador++
		}

		//validacion de valores y comandos iterados
		if strings.Compare(comandoingresado, "mkfs") == 0 {
			fmt.Println("Encontro: " + comandoingresado)
			comandoingresado = ""
			contador++
		} else if strings.Compare(comandoingresado, ">id=") == 0 {
			fmt.Println("Encontro :" + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			flag_id = true
			comandoingresado = ""
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_id += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor: " + valor_id)
			cadenaf += "Valor: " + valor_id + " \n"
		} else if strings.Compare(comandoingresado, ">type=") == 0 {
			fmt.Println("Encontro: " + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			comandoingresado = ""
			flag_type = true
			valor_type = ""

			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_type += strings.ToLower(linecomand[contador])
					contador++
				}
			}
			fmt.Println("Valor: " + valor_type)
			cadenaf += "Valor: " + valor_type + " \n"
		}
	}

	if flag_id {
		fmt.Println(valor_id)
		if flag_type {
			fmt.Println(valor_type)
			//aqui mano a cambiar el .formati
			flag_format = formateoEXT(valor_id)
			if flag_format {
				fmt.Println("Se genero de manera correcta el formateo EXT2")
				cadenaf += "Se genero de manera correcta el formateo EXT2" + " \n"
			} else {
				fmt.Println("ID no se encuentra montado por lo que no se puede generar EXT2")
				cadenaf += "ID no se encuentra montado por lo que no se puede generar EXT2" + " \n"
			}
		} else {
			fmt.Println("se deja valor por defecto: " + valor_type)
			//aqui mano a cambiar el .formati
			flag_format = formateoEXT(valor_id)
			if flag_format {
				fmt.Println("Se genero de manera correcta el formateo EXT2")
				cadenaf += "Se genero de manera correcta el formateo EXT2" + " \n"
			} else {
				fmt.Println("ID no se encuentra montado por lo que no se puede generar EXT2")
				cadenaf += "ID no se encuentra montado por lo que no se puede generar EXT2" + " \n"
			}
		}
	} else {
		fmt.Println("NO cuenta con comando Oblogatorio.. ")
		cadenaf += "NO cuenta con comando Oblogatorio.." + " \n"
	}
}