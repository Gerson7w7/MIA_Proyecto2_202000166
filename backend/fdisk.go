package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

//Metodo para la extraccion de datos:
func AnalisisFdisk(comando string) {
	var linecomand [100]string
	nuevocomando := strings.Split(comando, "")
	comandoingresado := ""
	//copio a array linecomando para mejor manejo
	copy(linecomand[:], nuevocomando[:])

	//banderas
	flag_size := false //obligatorio
	flag_unit := false
	flag_path := false //obligatorio
	flag_type := false
	flag_fit := false
	flag_name := false //obligatorio

	//valores almacenados
	valor_size := ""
	valor_unit := "k"
	valor_path := ""
	valor_type := "p"
	valor_fit := "wf"
	valor_name := ""

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
		if strings.Compare(comandoingresado, "fdisk") == 0 {
			fmt.Println("Encontro: " + comandoingresado)
			comandoingresado = ""
			contador++
		} else if strings.Compare(comandoingresado, ">size=") == 0 {
			fmt.Println("Encontro :" + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			flag_size = true
			comandoingresado = ""
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_size += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor: " + valor_size)
			cadenaf += "Valor: " + valor_size + " \n"
		} else if strings.Compare(comandoingresado, ">unit=") == 0 {
			fmt.Println("Encontro :" + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			comandoingresado = ""
			flag_unit = true
			//la asignacion es directa
			valor_unit = strings.ToLower(linecomand[contador])
			contador++
			fmt.Println("Valor: " + valor_unit)
			cadenaf += "Valor: " + valor_unit + " \n"
		} else if strings.Compare(comandoingresado, ">path=") == 0 {
			fmt.Println("Encontro :" + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			comandoingresado = ""
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
			fmt.Println("Valor: " + valor_path)
			cadenaf += "Valor: " + valor_path + " \n"
		} else if strings.Compare(comandoingresado, ">type=") == 0 {
			fmt.Println("Encontro: " + comandoingresado)
			comandoingresado = ""
			flag_type = true
			valor_type = ""
			//la asignacion es directa
			valor_type += strings.ToLower(linecomand[contador])
			contador++
			fmt.Println("Valor: " + valor_type)
		} else if strings.Compare(comandoingresado, ">fit=") == 0 {
			fmt.Println("Encontro: " + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			comandoingresado = ""
			flag_fit = true
			valor_fit = ""
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 {
					contador++
					break
				} else {
					valor_fit += strings.ToLower(linecomand[contador])
					contador++
				}
			}
			fmt.Println("Valor: " + valor_fit)
			cadenaf += "Valor: " + valor_fit + " \n"
		} else if strings.Compare(comandoingresado, ">name=") == 0 {
			fmt.Println("Encontro: " + comandoingresado)
			cadenaf += "Encontro: " + comandoingresado + " \n"
			comandoingresado = ""
			flag_name = true
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_name += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor: " + valor_name)
			cadenaf += "Valor: " + valor_name + " \n"
		}
	}
	//proceso para crear particiones
	crearParticion(flag_size, flag_unit, flag_path, flag_type, flag_fit, flag_name, valor_size, valor_unit, valor_path, valor_type, valor_fit, valor_name)
}

//validaciones para crear particion
func crearParticion(_flag_size bool, _flag_unit bool, _flag_path bool, _flag_type bool, _flag_fit bool, _flag_name bool, _size string, _unit string, _path string, _type string, _fit string, _name string) {
	//validaciones para crear particion
	size_entero := 0
	var auxsize_bytes int64
	var auxtype byte
	var auxfit string
	if _flag_size == true {
		size_entero, _ = strconv.Atoi(_size)
		if size_entero > 0 {
			fmt.Println("Tamaño aceptado")
			cadenaf += "Tamaño aceptado" + " \n"
		} else {
			fmt.Println("Error-> Tamaño invalido")
			cadenaf += "Error-> Tamaño invalido" + " \n"
			return
		}
	} else {
		fmt.Println("Error-> No viene el tamaño especificado")
		cadenaf += "Error-> No viene el tamaño especificado" + " \n"
		return
	}

	if _flag_unit == true {
		if strings.Compare(_unit, "b") == 0 {
			auxsize_bytes = int64(size_entero)
		} else if strings.Compare(_unit, "k") == 0 {
			auxsize_bytes = int64(size_entero) * 1024
		} else if strings.Compare(_unit, "m") == 0 {
			auxsize_bytes = int64(size_entero) * 1024 * 1024
		} else {
			fmt.Println("Error -> Unit invalido")
			cadenaf += "Error -> Unit invalido" + " \n"
		}
	} else {
		auxsize_bytes = int64(size_entero) * 1024
	}

	if _flag_type == true {
		if strings.Compare(_type, "p") == 0 {
			auxtype = 'p'
		} else if strings.Compare(_type, "e") == 0 {
			auxtype = 'e'
		} else if strings.Compare(_type, "l") == 0 {
			auxtype = 'l'
		} else {
			fmt.Println("Error -> Type invalido")
			cadenaf += "Error -> Type invalido" + " \n"
		}
	} else {
		_flag_type = true //aqui si no esta especificado osea que viene falso del analisis de parametrso
		auxtype = 'p'
	}

	if _flag_fit == true {
		if strings.Compare(_fit, "bf") == 0 {
			auxfit = "bf"
		} else if strings.Compare(_fit, "ff") == 0 {
			auxfit = "ff"
		} else if strings.Compare(_fit, "wf") == 0 {
			auxfit = "wf"
		} else {
			fmt.Println("Error-> Fit invalido: ", _fit, " .tiene que ser: wf, ff, bf")
			cadenaf += "Error-> Fit invalido" + " \n"
		}
	} else {
		auxfit = "wf"
	}

	//funcion para validar si existe disco a particionar
	availableDisk := validacionDisco(_path)
	if availableDisk == true { // existe disco
		disco := Mbr{} // la estructura para el disco
		//setea valor anteriores
		// se empieza a leer el mbr del disco
		var size_mbr int64 = int64(unsafe.Sizeof(disco))
		file, err := os.OpenFile(_path, os.O_RDWR, 0644) //perimiso de leer y escribir
		if err != nil {
			cadenaf += "Error al abrir disco" + " \n"
			log.Fatal("Error al abrir disco", err)
		} else {
			file.Seek(0, 0)
			data := leerSiguienteBytes(file, size_mbr)
			buffer := bytes.NewBuffer(data)
			//aqui le decimos que el buffer se convierte a informacion normal y la guaramos en el Disco
			err = binary.Read(buffer, binary.BigEndian, &disco)
			if err != nil {
				cadenaf += "Error al leer disco" + " \n"
				log.Fatal("Error al leer disco", err)
			}
		}
		// se termina de leer
		cadenaf += "El disco existe, se procede a particionar" + " \n"
		fmt.Println("El disco existe, se procede a particionar")
		
		contp := contadorPrimaria(disco)
		//primero que tipo de particion se necesita
		if _flag_type == true {
			if strings.Compare(_type, "p") == 0 { //particion primaria
				fmt.Println("Contador particiona primaria")
				fmt.Println(contp)
				fmt.Println("Se solicita una particion primaria")
				fmt.Println("Nombre-> " + _name)
				band_name := validacionNombre(disco, _name)
				fmt.Println(band_name)
				if contp < 4 { // todavia hay espacio para una mas
					if !band_name { // nombre no se repite
						for i := 0; i < 4; i++ {
							if strings.Compare(string(disco.Partition[i].Part_type), "-") == 0 { //está disponible
								fmt.Println("Entra para particionar")
								var size_Total int64 = int64(disco.Mbr_tamano) //Esto es el tamaño total del Disco
								//Esto Son los espacios que ya tiene cada Disco si no solo es cero
								var size_Usado int64 = int64(disco.Partition[0].Part_size) + int64(disco.Partition[1].Part_size) + int64(disco.Partition[2].Part_size) + int64(disco.Partition[3].Part_size)
								//Espacion Total del Disco-Lo que Ocupa el MBR-El tamaño decada Disco
								var size_Disponible int64 = size_Total - int64(unsafe.Sizeof(disco)) - size_Usado
								//conversion
								intSize, _ := strconv.Atoi(_size) //Este el el tamaño que se desea asignar a la nueva particion
								//si size_Disponible es mayor size, hay espacio
								if size_Disponible > int64(intSize) {
									//se le asignan valores
									fmt.Println("Si hay espacio disponible")
									disco.Partition[i].Part_status = '1'
									disco.Partition[i].Part_type = auxtype
									copy(disco.Partition[i].Part_fit[:], []byte(auxfit))
									disco.Partition[i].Part_start = int64(unsafe.Sizeof(disco)) + size_Usado + 1
									disco.Partition[i].Part_size = auxsize_bytes
									copy(disco.Partition[i].Part_name[:], []byte(_name))
									// reescribir le mbr, para guardar cambios
									file.Seek(0, 0)
									var binariof1 bytes.Buffer
									binary.Write(&binariof1, binary.BigEndian, disco)
									escribirNextBytes(file, binariof1.Bytes())
									//se escribe la particion
									fmt.Println("Direccion a particioanr")
									fmt.Println(disco.Partition[i].Part_start)
									file.Seek(disco.Partition[i].Part_start, 0)
									var binariof2 bytes.Buffer
									binary.Write(&binariof2, binary.BigEndian, disco.Partition[i])
									escribirNextBytes(file, binariof2.Bytes())
									au := disco.Partition[i].Part_start
									//contenido := strconv.FormatInt(au, 64)
									agregarParticionMount(_path, _name, auxsize_bytes, _type, _fit, au, "0")
									file.Close() // se cierra el archivo

									//ESTO ES PARA AÑADIR LA PARTICION EN EL MONTAJE Y QUE SEA DINAMICO ESTA PARTE
									//path string, name string, _size string, _type string, _fit string, _start string, _next string)
									cadenaf += "Particion realizada correctamente " + string(i) + " \n"
									fmt.Println("Particion realizada correctamente " + string(i))
								} else {
									fmt.Println("Error-> Espacio insuficiente")
									cadenaf += "Error-> Espacio insuficiente" + " \n"
									return
								}
								break
							}
						}
					} else {
						fmt.Println("Error -> Se repite un nombre de la particion primaria")
						cadenaf += "Error -> Se repite un nombre de la particion primaria" + " \n"
						return
					}
				}
			} else if strings.Compare(_type, "e") == 0 { //particion extendida
				fmt.Println("Se solicita particion extendida")
				cont_ext := contadorExt(disco)
				fmt.Println("Contador particion extendida")
				fmt.Println(cont_ext)
				if cont_ext < 1 { //para validar que solo exista una extendida
					band_namee := validacionNombre(disco, _name)
					if !band_namee { //para validar nombre
						for i := 0; i < 4; i++ {
							if strings.Compare(string(disco.Partition[i].Part_type), "-") == 0 { //disponible un espacio para ext
								var size_Total int64 = int64(disco.Mbr_tamano)
								var size_Use int64 = int64(disco.Partition[0].Part_size) + int64(disco.Partition[1].Part_size) + int64(disco.Partition[2].Part_size) + int64(disco.Partition[3].Part_size)
								var size_Disp int64 = size_Total - int64(unsafe.Sizeof(disco)) - size_Use
								size_e, _ := strconv.Atoi(_size)

								if size_Disp > int64(size_e) {
									// se crea la estructura
									extendida := Ebr{}
									//se le setea valores predeterminados
									extendida.Part_status = '-'
									extendida.Part_next = -1
									extendida.Part_start = 0
									extendida.Part_size = 0

									//se setean datos al mbr
									disco.Partition[i].Part_status = '1'
									disco.Partition[i].Part_type = auxtype
									disco.Partition[i].Part_size = auxsize_bytes
									//var size_Usedespues int64 = int64(disco.Partition[0].Part_size) + int64(disco.Partition[1].Part_size) + int64(disco.Partition[2].Part_size) + int64(disco.Partition[3].Part_size)
									disco.Partition[i].Part_start = int64(unsafe.Sizeof(disco)) + size_Use + 1
									copy(disco.Partition[i].Part_fit[:], []byte(auxfit))
									copy(disco.Partition[i].Part_name[:], []byte(_name))

									//se reescribe el mbr
									file.Seek(0, 0)
									var binext1 bytes.Buffer
									binary.Write(&binext1, binary.BigEndian, disco)
									escribirNextBytes(file, binext1.Bytes())

									//se escribe la particion, y nos posicionamos en donde empieza la paticion que es es espacio Usado +1+mbrtamaño
									file.Seek(disco.Partition[i].Part_start, 0)
									var binext2 bytes.Buffer
									binary.Write(&binext2, binary.BigEndian, disco.Partition[i]) //aqui colocamos todo lo que Ocupara la Particion
									escribirNextBytes(file, binext2.Bytes())

									//se posiciona para escribir el ebr
									file.Seek(disco.Partition[i].Part_start+int64(unsafe.Sizeof(extendida))+1, 0)
									var binext3 bytes.Buffer
									binary.Write(&binext3, binary.BigEndian, extendida)
									escribirNextBytes(file, binext3.Bytes())

									au := disco.Partition[i].Part_start
									//contenido := strconv.FormatInt(au, 64)
									agregarParticionMount(_path, _name, auxsize_bytes, _type, _fit, au, "0")

									file.Close() // se cierra el archivo

									fmt.Println("Particion extendida creada correctamente")
									cadenaf += "Particion extendida creada correctamente" + " \n"
									break
								} else {
									fmt.Println("Error -> Espacio insuficiente para crear particion extendida")
									cadenaf += "Error -> Espacio insuficiente para crear particion extendida" + " \n"
									return
								}
							}
						}
					} else {
						fmt.Println("Error-> Nombre se repite de particion extendida")
						cadenaf += "Error-> Nombre se repite de particion extendida" + " \n"
					}
				} else {
					fmt.Println("Error-> Ya existe una particion extendida")
					cadenaf += "Error-> Ya existe una particion extendida" + " \n"
				}
			} else if strings.Compare(_type, "l") == 0 { //para crear particion logica
				fmt.Println("Se solicita particion logica")
				ebr := Ebr{}
				size_ebrl := int64(unsafe.Sizeof(ebr))
				counte := contadorLogica(disco, ebr, file)
				fmt.Println("Contador particion logica")
				fmt.Println(counte)
				if counte <= 24 {
					//para validar nombres en logicas
					flag_namel := validarNombreLogica(disco, ebr, file, _name)
					if !flag_namel {
						fmt.Println("No se repite nombre...")
						for i := 0; i < 4; i++ {
							if strings.Compare(string(disco.Partition[i].Part_type), "e") == 0 {
								fmt.Println("Se localiza una particion extendida...")
								file.Seek(disco.Partition[i].Part_start+int64(unsafe.Sizeof(ebr))+1, 0)
								datal := leerSiguienteBytes(file, size_ebrl)
								bufferl := bytes.NewBuffer(datal)
								err := binary.Read(bufferl, binary.BigEndian, &ebr)
								if err != nil {
									cadenaf += "Error al leer ebr" + " \n"
									log.Fatal("Error al leer ebr", err)
								}
								//validacion si es primero
								if ebr.Part_next == -1 {
									//siempre vendra lo de donde inicio +EBR+1
									ebr.Part_start = disco.Partition[i].Part_start + int64(unsafe.Sizeof(ebr)) + 1
									ebr.Part_status = '1'
									ebr.Part_size = auxsize_bytes
									ebr.Part_next = ebr.Part_start + ebr.Part_size + 1 // su siguiente es el otro que se crea abajo
									copy(ebr.Part_fit[:], []byte(auxfit))
									copy(ebr.Part_name[:], []byte(_name))

									//escribo el primer ebr
									file.Seek(ebr.Part_start, 0)
									var binl1 bytes.Buffer
									binary.Write(&binl1, binary.BigEndian, ebr)
									escribirNextBytes(file, binl1.Bytes())

									//nuevo ebr
									ebrnew := Ebr{}
									ebrnew.Part_next = -1
									ebrnew.Part_status = '-'
									ebrnew.Part_size = 0
									ebrnew.Part_start = 0
									//escribo el siguiente ebr
									file.Seek(ebr.Part_next, 0)
									var binl2 bytes.Buffer
									binary.Write(&binl2, binary.BigEndian, ebrnew)
									escribirNextBytes(file, binl2.Bytes())

									au := ebrnew.Part_start
									//contenido := strconv.FormatInt(au, 64)
									agregarParticionMount(_path, _name, auxsize_bytes, _type, _fit, au, "0")

									fmt.Println("Particion logica agregada correctamente")
									cadenaf += "Particion logica agregada correctamente" + " \n"

									file.Close() //se cierra el archivo

								} else { //no es primero, pero crean más ebr para ir estructurando la lista
									var pos int64 = ebr.Part_next
									var ant int64
									var size_eb int64 = int64(unsafe.Sizeof(ebr))
									//simula un while
									for pos != -1 {
										ant = pos
										file.Seek(pos, 0)
										data := leerSiguienteBytes(file, size_eb)
										buffer := bytes.NewBuffer(data)
										err := binary.Read(buffer, binary.BigEndian, &ebr)
										if err != nil {
											cadenaf += "Error al leer ebr" + " \n"
											log.Fatal("Error al leer ebr", err)
										}
										pos = ebr.Part_next
									}

									if counte <= 23 {
										ebr.Part_start = ant
										ebr.Part_status = '1'
										ebr.Part_size = auxsize_bytes
										ebr.Part_next = ebr.Part_start + ebr.Part_size + 1
										copy(ebr.Part_fit[:], []byte(auxfit))
										copy(ebr.Part_name[:], []byte(_name))

										//se escribe
										file.Seek(ebr.Part_start, 0)
										var binl3 bytes.Buffer
										binary.Write(&binl3, binary.BigEndian, ebr)
										escribirNextBytes(file, binl3.Bytes())
										fmt.Println("Particion logica agregada correctamente")
										cadenaf += "Particion logica agregada correctamente" + " \n"

										au := ebr.Part_start
										//contenido := strconv.FormatInt(au, 64)
										agregarParticionMount(_path, _name, auxsize_bytes, _type, _fit, au, "0")
									} else {
										if counte == 24 && ebr.Part_status == '0' {
											ebr.Part_start = ant
											ebr.Part_status = '1'
											ebr.Part_size = auxsize_bytes
											ebr.Part_next = -1
											copy(ebr.Part_fit[:], []byte(auxfit))
											copy(ebr.Part_name[:], []byte(_name))

											//se posiciona y se escribe
											file.Seek(ebr.Part_start, 0)
											var binl4 bytes.Buffer
											binary.Write(&binl4, binary.BigEndian, ebr)
											escribirNextBytes(file, binl4.Bytes())
											fmt.Println("Particion logica agregada correctamente")
											cadenaf += "Particion logica agregada correctamente" + " \n"

											au := ebr.Part_start
											//contenido := strconv.FormatInt(au, 64)
											agregarParticionMount(_path, _name, auxsize_bytes, _type, _fit, au, "0")

										} else {
											fmt.Println("Error -> Particiones logicas llego al maximo")
											cadenaf += "Error -> Particiones logicas llego al maximo" + " \n"
										}
									}
									if counte <= 23 {
										ebrnew := Ebr{}
										ebrnew.Part_next = -1
										ebrnew.Part_size = 0
										ebrnew.Part_start = 0
										ebrnew.Part_status = '0'
										file.Seek(ebr.Part_next, 0)
										var binl5 bytes.Buffer
										binary.Write(&binl5, binary.BigEndian, ebrnew)
										escribirNextBytes(file, binl5.Bytes())

									}
									file.Close() //se cierra el archivo
								}
								break
							} else {
								fmt.Println("Error-> Al parecer no hay particion extendida")
								cadenaf += "Error-> Al parecer no hay particion extendida" + " \n"
							}
						}
					} else {
						fmt.Println("Error-> Se repite nombre en particiones logicas")
						cadenaf += "Error-> Se repite nombre en particiones logicas" + " \n"
					}
				} else {
					fmt.Println("Error-> Se llego al maximo número de particiones logicas")
					cadenaf += "Error-> Se llego al maximo número de particiones logicas" + " \n"
				}

			} else {
				fmt.Println("Error-> Tipo invalido de particion")
				cadenaf += "Error-> Tipo invalido de particion" + " \n"
				return
			}
		} else {
			fmt.Println("Error-> Tipo de particion no especificado")
			cadenaf += "Error-> Tipo de particion no especificado" + " \n"
			return
		}
	} else { //no existe disco
		fmt.Println("Erro->  El disco no existe " + _path)
		cadenaf += "Error->  El disco no existe " + " \n"
		return
	}
}

func validacionDisco(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func leerSiguienteBytes(file *os.File, number int64) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal("Error al abrir byte ", err)
	}
	return bytes
}

//cuenta cuantas particiones primarias Hay
func contadorPrimaria(disco Mbr) int {
	cont := 0
	for i := 0; i < 4; i++ {
		if strings.Compare(string(disco.Partition[i].Part_type), "p") == 0 {
			cont++
		}
	}
	return cont
}

func contadorExt(disco Mbr) int {
	cont := 0
	for i := 0; i < 4; i++ {
		if strings.Compare(string(disco.Partition[i].Part_type), "e") == 0 {
			cont++
		}
	}
	return cont
}

//valido el nombre de la particion para que no se repite
func validacionNombre(disco Mbr, name string) bool {
	fmt.Println("nombre a buscar " + name)
	bandera := false
	for i := 0; i < 4; i++ {
		fmt.Println("-->" + string(disco.Partition[i].Part_name[:]) + "<--")
		auxname := string(disco.Partition[i].Part_name[:len(name)]) //como el nombre se guarda en un arreglo debe de reconstruirse
		fmt.Println(len(auxname))
		if strings.Compare(auxname, name) == 0 {
			fmt.Println("Es igual")
			bandera = true
			break
		}
		fmt.Println("No es igual")
	}
	return bandera
}

func escribirNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func contadorLogica(disco Mbr, ebr Ebr, file *os.File) int {
	fmt.Println("Entra en contador logica...")
	contebr := 0
	var next int64
	var tam int64 = int64(unsafe.Sizeof(ebr))
	for i := 0; i < 4; i++ {
		if strings.Compare(string(disco.Partition[i].Part_type), "e") == 0 {
			//nos posciionames en Particio+el espacionEBR+1
			file.Seek(disco.Partition[i].Part_start+int64(unsafe.Sizeof(ebr))+1, 0)
			data := leerSiguienteBytes(file, tam)
			buffer := bytes.NewBuffer(data)
			err := binary.Read(buffer, binary.BigEndian, &ebr)
			if err != nil {
				fmt.Println("Error al leer disco")
				log.Fatal(err)
			}
			fmt.Println("Paso 1")
			fmt.Println(ebr.Part_next)
			fmt.Println(ebr.Part_size)
			fmt.Println(ebr.Part_start)
			fmt.Println(ebr.Part_status)
			next = int64(ebr.Part_next)

			//simula un while
			for next != -1 {
				tame := int64(unsafe.Sizeof(ebr))
				file.Seek(next, 0) //me posiciono
				date := leerSiguienteBytes(file, tame)
				buffere := bytes.NewBuffer(date)
				erre := binary.Read(buffere, binary.BigEndian, &ebr)
				if erre != nil {
					log.Fatal("Error al leer ebr", erre)
				}
				next = ebr.Part_next
				contebr++
			}
			contebr++
		}
	}
	return contebr
}

func validarNombreLogica(disco Mbr, ebr Ebr, file *os.File, name string) bool {
	var flag_name bool = false
	var taml int64 = int64(unsafe.Sizeof(ebr))
	next := 0
	for i := 0; i < 4; i++ {
		if strings.Compare(string(disco.Partition[i].Part_type), "e") == 0 {
			file.Seek(disco.Partition[i].Part_start+int64(unsafe.Sizeof(ebr))+1, 0)
			datal := leerSiguienteBytes(file, taml)
			bufferl := bytes.NewBuffer(datal)
			errl := binary.Read(bufferl, binary.BigEndian, &ebr)
			if errl != nil {
				log.Fatal("Error al leer el disco", errl)
			}
			next = int(ebr.Part_next)
			//simula un while
			for next != -1 {
				if strings.Compare(string(ebr.Part_name[:len(name)]), name) == 0 {
					flag_name = true
					break
				}
				file.Seek(int64(next), 0)
				datal := leerSiguienteBytes(file, taml)
				bufferl := bytes.NewBuffer(datal)
				errl := binary.Read(bufferl, binary.BigEndian, &ebr)
				if errl != nil {
					log.Fatal("Error al leer el disco", errl)
				}
				next = int(ebr.Part_next)
			}
			if strings.Compare(string(ebr.Part_name[:]), name) == 0 {
				flag_name = true
			}
		}
	}
	return flag_name
}