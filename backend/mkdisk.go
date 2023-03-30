package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func AnalisiMkdisk(comando string) {
	fmt.Println("Entramos al MKDISK")
	var linecomand [200]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""
	// se realiza una copia del array para manejo
	copy(linecomand[:], newcomando[:])

	//banderas
	flag_size := false
	flag_fit := false
	flag_unit := false
	flag_path := false

	//valores
	valor_size := ""
	valor_fit := "ff"
	valor_unit := "m"
	valor_path := ""

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
		if strings.Compare(lineacomando, "mkdisk") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
			//aqui limpio la linea comando y el contador sige en aumento para validar los demas caracteres.
		} else if strings.Compare(lineacomando, ">size=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n" //Esto es lo que se mostrara en el Front
			lineacomando = ""                              //se limpia linea
			flag_size = true                               //se activa bandera size	 ya que es comando obligatorio
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					//al entrar aqui validamos q si es espacio vacio o salto de linea se sale del while o for
					contador++
					break
				} else {
					valor_size += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor: " + valor_size)
			cadenaf += "Valor: " + valor_size + " \n"
		} else if strings.Compare(lineacomando, ">fit=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n"
			lineacomando = ""
			valor_fit = ""
			flag_fit = true
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_fit += strings.ToLower(linecomand[contador])
					contador++
				}
			}
			fmt.Println("Valor: " + valor_fit)
			cadenaf += "Valor: " + valor_fit + " \n"
		} else if strings.Compare(lineacomando, ">unit=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n"
			lineacomando = ""
			valor_unit = ""
			flag_unit = true
			//directo
			valor_unit = strings.ToLower(linecomand[contador])
			contador++
			fmt.Println("Valor: " + valor_unit)
			cadenaf += "Valor: " + valor_unit + " \n"
		} else if strings.Compare(lineacomando, ">path=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			cadenaf += "Encontro: " + lineacomando + " \n"
			lineacomando = ""
			flag_path = true
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
							valor_path += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_path += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor : " + valor_path)
			cadenaf += "Valor : " + valor_path + " \n"
		}
	}
	//---------------PROCESO CREACION DE DISCOS-------------------------
	fmt.Println("Inicio de proceso")
	contadorDiagonal := 0
	for _, ele := range valor_path { // se cuenta cuantas diagonales hay para directorio
		if strings.Compare(string(ele), "/") == 0 {
			contadorDiagonal++
		}
	}
	valor_directorio := ""
	auxContador := 0
	for _, ele := range valor_path { // se obtiene solo directorio
		if strings.Compare(string(ele), "/") == 0 {
			valor_directorio += string(ele)
			auxContador++
			if contadorDiagonal == auxContador {
				break
			}
		} else {
			valor_directorio += string(ele)
		}
	}
	fmt.Println("El directorio a validar: ", valor_directorio)

	flag_directorio := validacionDirectorio(valor_directorio) // funcion que valida si existe el directorio
	flag_disco := validacionArchivo(valor_path)               //funcion que valida si existe el disco
	//validacion directorio
	if flag_directorio == true { // existe el directorio
		fmt.Println("¡Existe Directorio!")
		cadenaf += "¡Existe Directorio! \n"
		if flag_disco == true {
			fmt.Println("Error-> ¡El disco ya existe con ese nombre!")
			cadenaf += "Error-> ¡El disco ya existe con ese nombre! \n"
		} else {
			fmt.Println("¡El disco no existe, se procede con la creación!")
			cadenaf += "¡El disco no existe, se procede con la creación! \n"
			crearDisco(flag_size, flag_unit, flag_path, flag_fit, valor_size, valor_path, valor_unit, valor_fit)
		}
	} else { // no existe el directorio
		fmt.Println("¡Directorio no existe!")
		cadenaf += "¡Directorio no existe! \n"
		SoloCrearDirectorio(valor_directorio)
		fmt.Println("¡Directorio creado exitosamente!")
		cadenaf += "¡Directorio creado exitosamente! \n"
		crearDisco(flag_size, flag_unit, flag_path, flag_fit, valor_size, valor_path, valor_unit, valor_fit)
	}
	//para leer el disco si se creo correctamente
	LeerDisco(valor_path)
}

func SoloCrearDirectorio(valor_directorio string) {
	fmt.Println("EL directorio a crear: ", valor_directorio, "------")
	merr := os.Mkdir(valor_directorio, 0777) //se agrega los permisos para la creacion del directorio estaba 0755
	//func Mkdir(name string, perm FileMode)
	if merr != nil { //si hay errores no se genera el Disco
		fmt.Println("Error al crear directorio -> " + valor_directorio)
		log.Fatal(merr)
	}
}

func validacionDirectorio(directorio string) bool {
	if _, err := os.Stat(directorio); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func validacionArchivo(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func crearDisco(_flag_size bool, _flag_unit bool, _flag_path bool, _flag_fit bool, size string, path string, unit string, fit string) {
	disco := Mbr{}   // se crea la estructura de mbr
	var tamano int64 // tipo i64 es el tipo int de mayor rango

	if _flag_size == true {
		_tamano, err := strconv.Atoi(size) //realiza la conversion a entero
		if err != nil {
			log.Fatal(err)
		}
		if _tamano > 0 {
			fmt.Println("¡Tamaño disco valido!")
			cadenaf += "¡Tamaño disco valido! \n"
			tamano = int64(_tamano) //aqui se realiza a base 64
		} else {
			fmt.Println("Error-> ¡Tamaño de disco no valido!")
			cadenaf += "Error-> ¡Tamaño de disco no valido! \n"
		}
	}

	if _flag_fit == true {
		if strings.Compare(fit, "bf") == 0 {
			copy(disco.Dsk_fit[:], fit)
		} else if strings.Compare(fit, "ff") == 0 {
			copy(disco.Dsk_fit[:], fit)
		} else if strings.Compare(fit, "wf") == 0 {
			copy(disco.Dsk_fit[:], fit)
		} else {
			fmt.Println("Error-> ¡Valor invalido de fit!")
			cadenaf += "Error-> ¡Valor invalido de fit! \n"
		}
	} else {
		copy(disco.Dsk_fit[:], "ff") // si no es especificado es ff
	}

	if _flag_unit == true {
		fmt.Println("Unit " + unit)
		if strings.Compare(unit, "k") == 0 { // si es kilobytes
			disco.Mbr_tamano = int64(tamano) * 1024
		} else if strings.Compare(unit, "m") == 0 { // si es megabytes
			disco.Mbr_tamano = int64(tamano) * 1024 * 1024
		} else {
			fmt.Println("Error-> ¡Valor de unit invalido!")
			cadenaf += "Error-> ¡Valor de unit invalido! \n"
		}
	} else {
		disco.Mbr_tamano = int64(tamano) * 1024 * 1024 // si no es especificado es megabytes
	}

	if _flag_path == true {
		disco.Mbr_dsk_signature = int64(rand.Intn(100))
		//iniciando valores de particion
		for i := 0; i < 4; i++ {
			disco.Partition[i].Part_status = '0' //valor en ascci este significa caracter nulo
			disco.Partition[i].Part_type = '-'
			disco.Partition[i].Part_start = 0
			disco.Partition[i].Part_size = 0

		}
		file, err := os.Create(path) //CREO EL ARCHIVO COMPLETO
		//se declara con 2 variables ya que el metodo o.create genera dos valores si o si
		if err != nil { //nil significa cero errores al generar algun proceso de lectura
			log.Fatal(err)
		}

		//colocar el primer cero
		var vacio int8 = 0
		cero := &vacio
		var binario1 bytes.Buffer // nose para que sirve
		binary.Write(&binario1, binary.BigEndian, cero)
		writeNextBytes(file, binario1.Bytes())

		//posicionado en la ultima posicion
		var binario2 bytes.Buffer
		file.Seek(disco.Mbr_tamano, 0)                  //aqui indicamos con 0 que nos ubicamos en el origen del archivo y que apartir de aqui iremos al .Mbr tamaño
		binary.Write(&binario2, binary.BigEndian, cero) // esto es como el inserta el valor de cero en binario
		writeNextBytes(file, binario2.Bytes())

		//se genera la fecha
		fecha := time.Now()
		fechaSep := strings.Split(fecha.String(), "")
		fechareal := ""
		for i := 0; i < 16; i++ {
			fechareal += fechaSep[i]
		}
		copy(disco.Mbr_fecha_creacion[:], fechareal)
		//para empezar a escribir el mbr desde el principio
		file.Seek(0, 0) //pues entiendo que llegara del rango 0INicial al --------------------0 final
		//escribir el mbr
		var binario3 bytes.Buffer
		binary.Write(&binario3, binary.BigEndian, disco)
		writeNextBytes(file, binario3.Bytes())
		file.Close()
	}
}

func writeNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func LeerDisco(path string) {
	fmt.Println("Path a abrir")
	fmt.Println(path)
	m := Mbr{}
	var tam_mbr int64 = int64(unsafe.Sizeof(m))
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error al abrir disco", err)
	} else {
		file.Seek(0, 0)
		data := leerSiguienteByte(file, tam_mbr)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &m)
		if err != nil {
			fmt.Println("Error al leer disco")
			log.Fatal(err)
		}
	}
	fmt.Println("------------LEER--------------")
	tam := int64(m.Mbr_tamano)
	fmt.Println("Tamaño: ", tam)
	dsk := int64(m.Mbr_dsk_signature)
	fmt.Println("Signature: ", dsk)
	fit := string(m.Dsk_fit[:])
	fmt.Println("Fit: ", fit)
	fech := string(m.Mbr_fecha_creacion[:])
	fmt.Println("Fecha: ", fech)
	fmt.Println("------------------------------")
	file.Close()
}

func leerSiguienteByte(file *os.File, number int64) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal("Error al abrir byte", err)
	}
	return bytes
}