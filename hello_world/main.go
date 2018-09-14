package main

import "fmt"

const hello string = "Hola %s, %s. Bienvenido!!!\n"
const helloworld = "Hola Mundo"

func main() {

	name, lastname := getName()
	var number = sum(45, 62)

	//Manera de declarar multiples variables seguidas sin escribir "var" cade vez.
	var (
		a        = 1
		b string = "Algo"
		c        = true
	)

	fmt.Printf(hello, name, lastname)
	fmt.Println(helloworld)
	fmt.Println(number, a, b, c)
}

func getName() (string, string) {
	//Manera estatica de declarar una variable. El tipo de dato se puede omitir y go lo deducura.
	var name string = "Nombre"
	//Version corta de declaracion de variables. ":=" solo es usado para declaracion, no para asignacion ("=")
	lastname := "Apellido"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s %s", &name, &lastname)
	return lastname, name

}

func sum(a int, b int) int {
	return a + b
}
