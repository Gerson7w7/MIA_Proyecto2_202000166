package main

import (
	"fmt"
	"strings"
	"container/list"
)

// variables globales
var cadenaf string = ""
var imagenFinalRep string = ""
var repVali = false
var ReporteFInal = ""
var Users = list.New()
var Groups = list.New() 

func AnalizadorComando(comando string) {
	lineacomando := "" // donde se guarda el primer comando
	contador := 0      //contador general para recorrer el comando
	comandosep := strings.Split(comando, "")

	//comprube si viene vacio el comando
	if strings.Compare(comandosep[0], "\n") == 0 {
		fmt.Println("Salto de linea")
		cadenaf += "Salto de linea \n" // para retornar
	} else {
		//simula un while
		for (strings.Compare(comandosep[contador], "\n") != 0) && (strings.Compare(comandosep[contador], "golang\000") != 0) { // si no viene vacio -> \n
			if strings.Compare(comandosep[contador], " ") == 0 { // si viene espacio
				// aqui solo valida el comando no sus atributos en este caso exec, mkdisk etc
				break
			} else {
				lineacomando += strings.ToLower(comandosep[contador]) // va concatenando cada char del comando
				contador++
			}
		}
	}

	//aqui ya valido la ifnoramcion
	if strings.Compare(lineacomando, "execute") == 0 {
		//AnalisisExec(comando)
		fmt.Println("		Se encontro el execute: ", comando)
		cadenaf += "Se encontro el Exec: " + string(comando) + " \n"
	} else if strings.Compare(lineacomando, "mkdisk") == 0 {
		AnalisiMkdisk(comando)
		fmt.Println("		Se encontro el mkdisk: ", comando)
		cadenaf += "Se encontro el mkdisk: " + string(comando) + " \n"
	} else if strings.Compare(lineacomando, "rmdisk") == 0 {
		AnalisisRmdisk(comando)
		fmt.Println("		Se encontro el rmdisk: ", comando)
		cadenaf += "Se encontro el rmdisk: " + string(comando) + " \n"
	} else if strings.Compare(lineacomando, "fdisk") == 0 {
		AnalisisFdisk(comando)
		fmt.Println("		Se encontro el fdisk: ", comando)
		cadenaf += "Se encontro el fdisk: " + string(comando) + " \n"
	} else if strings.Compare(lineacomando, "mount") == 0 {
		AnalisisMount(comando)
		MonstrarMount()
		fmt.Println("		Se encontro el mount: ", comando)
		cadenaf += "Se encontro el mount: " + string(comando) + " \n"
	} else if strings.Compare(lineacomando, "mkfs") == 0 {
		//fmt.Println("Realizando formateo EXT2")
		analisisMkfsv2(comando)
		fmt.Println("		Se encontro el mkfs: ", comando)
		cadenaf += "Se encontro el mkfs: " + string(comando) + "\n"
	} else if strings.Compare(lineacomando, "rep") == 0 {
		AnalsisRep(comando)
		fmt.Println("		Se encontro el rep: ", comando)
		cadenaf += "Se encontro el rep: " + string(comando) + " \n"
	} else if strings.Compare(lineacomando, "pause") == 0 {
		//pause
		fmt.Println("		Pause.............")
		cadenaf += "Se encontro el Pause:......... \n"
	} else if strings.Compare(lineacomando, "mkgrp") == 0 {
		//mkgrp
		fmt.Println("		Se encontro el mkgrp: ", comando)
		cadenaf += "Se encontro el mkgrp:......... \n"
		AnalisiMkgrp(comando)
	} else if strings.Compare(lineacomando, "rmgrp") == 0 {
		//rmgrp
		fmt.Println("		Se encontro el rmgrp: ", comando)
		cadenaf += "Se encontro el rmgrp:......... \n"
		AnalisiRmgrp(comando)
	} else if strings.Compare(lineacomando, "mkusr") == 0 {
		//mkusr
		fmt.Println("		Se encontro el mkusr: ", comando)
		cadenaf += "Se encontro el mkusr:......... \n"
		AnalisiMkusr(comando)
	} else if strings.Compare(lineacomando, "rmusr") == 0 {
		//rmuser
		fmt.Println("		Se encontro el rmuser: ", comando)
		cadenaf += "Se encontro el rmusr:......... \n"
		AnalisiRmusr(comando)
	} else {
		fmt.Println("Se encontro un comentario: " + string(comando) + "\n")
		cadenaf += "Se encontro un comentario: " + string(comando) + "\n"
	}
}