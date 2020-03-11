package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Alumns struct {
	Nombre string
	Cal1   int
	Cal2   int
	Cal3   int
}

var Alumnos []Alumns

func Esperar() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("presione enter para continuar")
	reader.ReadString('\n')
}

func Clear() {
	// print("\033[H\033[2J")
	// fmt.Print("\x0c")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Mostrar(i int) {

	fmt.Printf("ID: %d \n", i)
	fmt.Printf("nombre: %s \n", Alumnos[i].Nombre)
	fmt.Printf("calificacion 1: %d \n", Alumnos[i].Cal1)
	fmt.Printf("calificacion 2: %d \n", Alumnos[i].Cal2)
	fmt.Printf("calificacion 3: %d \n", Alumnos[i].Cal3)
}

func Lista() {
	Clear()
	for i := 0; i < len(Alumnos); i++ {
		fmt.Println(strings.Repeat("*", 10))
		Mostrar(i)
		fmt.Println(strings.Repeat("*", 10))
	}
	Esperar()
}

func GuardadInfo() {
	Clear()
	hello2 := ""
	for i := 0; i < len(Alumnos); i++ {

		cal1Parsed := strconv.Itoa(Alumnos[i].Cal1)
		cal2Parsed := strconv.Itoa(Alumnos[i].Cal2)
		cal3Parsed := strconv.Itoa(Alumnos[i].Cal3)
		hello2 += Alumnos[i].Nombre + " " + cal1Parsed + " " + cal2Parsed + " " + cal3Parsed + "\n"
	}
	f, _ := os.Create("test.txt")
	defer f.Close()
	f.WriteString(hello2)

	fmt.Println("Escrita con Exito!!")
	Esperar()
}

func LeerInfo() {
	Clear()
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("The file contains: " + string(data))
	Esperar()
}
func Capturar() {
	Clear()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ingrese el nombre del alumno:")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Println("ingrese calificacion 1:")
	cal1, _ := reader.ReadString('\n')
	cal1 = strings.TrimSpace(cal1)

	fmt.Println("ingrese calificacion 2:")
	cal2, _ := reader.ReadString('\n')
	cal2 = strings.TrimSpace(cal2)

	fmt.Println("ingrese calificacion 3:")
	cal3, _ := reader.ReadString('\n')
	cal3 = strings.TrimSpace(cal3)

	cal1Parsed, _ := strconv.Atoi(cal1)
	cal2Parsed, _ := strconv.Atoi(cal2)
	cal3Parsed, _ := strconv.Atoi(cal3)
	var Alumn = Alumns{
		Nombre: nombre,
		Cal1:   cal1Parsed,
		Cal2:   cal2Parsed,
		Cal3:   cal3Parsed,
	}
	Alumnos = append(Alumnos, Alumn)
	Esperar()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	flag := false
	for flag == false {
		Clear()
		fmt.Println("--MENU--")
		fmt.Println("1)Capturar")
		fmt.Println("2)Guardar Informacion")
		fmt.Println("3)Abrir archivo")
		fmt.Println("4)Salir")
		fmt.Println("5)lista")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch option {
		case "4":
			flag = true
		case "1":
			Capturar()
		case "2":
			GuardadInfo()
		case "3":
			LeerInfo()
		case "5":
			Lista()

		}

	}
}
