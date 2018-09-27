package main

import "github.com/contreras117/gocurso/interfaces/interpck"

func main() {

	testSuscribe()
}

func testSuscribe() {
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
