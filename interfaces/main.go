package main

import (
	"fmt"

	"github.com/contreras117/gocurso/interfaces/interpck"
)

func main() {

	testSuscribe()
}

func testSuscribe() {

	//La palabra reservada defer permite ejecutar una funcion al terminar la ejecucion del funcion detro de la que se encuentra.
	//Puede ir en cualquier lugar dentro de la funcion, ya se al inicio, en el medio, al final, etc.
	defer deferTest()
	curso := interfaces.Course{Name: "Progrmacion con go", Slug: "go", Skills: []string{"Backend", "Servidores", "Programacion"}}
	curso2 := interfaces.Course{Name: "Bases de datos", Slug: "BDD", Skills: []string{"Backend", "Servidores"}}

	carrera := new(interfaces.Career)
	carrera.Name = "Backend con Go"
	carrera.Courses = []interfaces.Course{curso, curso2}

	callSuscribe(curso)
	callSuscribe(carrera)
}

func callSuscribe(p interfaces.Platzi) {
	p.Suscribe("Daniel")
}

func deferTest() {
	fmt.Println("La funcion testSuscribe ha terminado")
}
