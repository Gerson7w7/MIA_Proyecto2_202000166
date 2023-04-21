package main

import (
	"fmt"
	"strings"
)

func AnalisiMkusr(comando string) {
	fmt.Println("Entramos al mkusr")
	var linecomand [200]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""
	// se realiza una copia del array para manejo
	copy(linecomand[:], newcomando[:])

	//banderas
	flag_user := false
	flag_pwd := false
	flag_grp := false

	//valores
	valor_user := ""
	valor_pwd := ""
	valor_grp := ""

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
		if strings.Compare(lineacomando, "mkusr") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
			//aqui limpio la linea comando y el contador sige en aumento para validar los demas caracteres.
		} else if strings.Compare(lineacomando, ">user=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n" //Esto es lo que se mostrara en el Front
			lineacomando = ""                              //se limpia linea
			flag_user = true                               //se activa bandera size	 ya que es comando obligatorio
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
							valor_user += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_user += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_user)
			cadenaf += "Valor: " + valor_user + " \n"
		} else if strings.Compare(lineacomando, ">pwd=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n"
			lineacomando = ""
			valor_pwd = ""
			flag_pwd = true
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
							valor_pwd += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_pwd += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_pwd)
			cadenaf += "Valor: " + valor_pwd + " \n"
		} else if strings.Compare(lineacomando, ">grp=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n"
			lineacomando = ""
			valor_grp = ""
			flag_grp = true
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
							valor_grp += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_grp += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_grp)
			cadenaf += "Valor: " + valor_grp + " \n"
		}
	}
	//---------------PROCESO CREACION DE USUARIOS-------------------------
	if (flag_user && flag_pwd && flag_grp) {
		fmt.Println("Inicio de proceso")
		for e := Users.Front(); e != nil; e = e.Next() {
			user := e.Value.(User)
			if valor_user == user.Name {
				fmt.Println("Error-> El nombre de usuario ya existe, elija otro diferente: " + valor_user)
				cadenaf += "Error-> El nombre de usuario ya existe, elija otro diferente: " + valor_user + "\n"
				return;
			}
		}
		for e := Groups.Front(); e != nil; e = e.Next() {
			group := e.Value
			if group == valor_grp {
				user := User{ Name: valor_user, Pwd: valor_pwd, Grp: valor_grp }
				Users.PushFront(user)
				fmt.Println("Usuario agregado con éxito: " + valor_user)
				cadenaf += "Usuario agregado con éxito: " + valor_user + " \n"
				for i := Users.Front(); i != nil; i = i.Next() {
					fmt.Println(i)
				}
				return;
			}
		}
		fmt.Println("Error-> Se necesita un grupo existente para poder crear el usuario: " + valor_grp)
		cadenaf += "Error-> Se necesita un grupo existente para poder crear el usuario: " + valor_grp + "\n"
	} else {
		fmt.Println("Error-> Se necesita usuario, contrasenia y grupo para poder crear un nuevo usario")
		cadenaf += "Error-> Se necesita usuario, contrasenia y grupo para poder crear un nuevo usario \n"
	}
}

func AnalisiRmusr(comando string) {
	fmt.Println("Entramos al rmusr")
	var linecomand [200]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""
	// se realiza una copia del array para manejo
	copy(linecomand[:], newcomando[:])

	//banderas
	flag_user := false

	//valores
	valor_user := ""

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
		if strings.Compare(lineacomando, "rmusr") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
			//aqui limpio la linea comando y el contador sige en aumento para validar los demas caracteres.
		} else if strings.Compare(lineacomando, ">user=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n" //Esto es lo que se mostrara en el Front
			lineacomando = ""                              //se limpia linea
			flag_user = true                               //se activa bandera size	 ya que es comando obligatorio
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
							valor_user += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_user += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_user)
			cadenaf += "Valor: " + valor_user + " \n"
		} 
	}
	//---------------PROCESO ELIMINACION DE USUARIOS-------------------------
	if (flag_user) {
		fmt.Println("Inicio de proceso")
		for e := Users.Front(); e != nil; e = e.Next() {
			user := e.Value.(User)
			if (valor_user == user.Name) {
				Users.Remove(e)
				fmt.Println("Usuario eliminado con éxito: " + valor_user)
				cadenaf += "Usuario eliminado con éxito: " + valor_user + " \n"
				for e := Users.Front(); e != nil; e = e.Next() {
					fmt.Println(e)
				}
				return;
			}
		}
		fmt.Println("Usuario no encontrado, verifique el nombre: " + valor_user)
		cadenaf += "Usuario no encontrado, verifique el nombre: " + valor_user + " \n"
	} else {
		fmt.Println("Error-> Se necesita el nombre de usuario para poder eliminar un usario")
		cadenaf += "Error-> Se necesita el nombre de usuario para poder eliminar un usario \n"
	}
}