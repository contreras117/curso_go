package main

import (
	"fmt"

	. "github.com/contreras117/gocurso/structs/course"
)

//MiDato - Con la palabra reservada type podemos crear nuevos tipos de datos a partir de otros.
type MiDato int

//OtroCurso Poniendo un struct dentro de la declararion de otro, este ultimo hereda todos los campos del primero, algo muy parecido a herencia.
type OtroCurso struct {
	Course
}

//Career es un struct que en sus valores internos tiene un arreglo de otro struct
type Career struct {
	courses []Course
}

func main() {

	//Esta es la manera larga y descriptiva de inicializar una instancia de struct, asignando sus valores desde un inicio.
	curso := Course{Name: "Progrmacion con go", Slug: "go", Skills: []string{"Backend", "Servidores", "Programacion"}}
	//Con la funcion new podemos inicializar un struct para despues asignarle los valores de sus elemetos. La diferencia es que la funcion new crea la variable como una apuntador del struct, en este caso *Course.
	curso2 := new(Course)
	curso2.Name = "Bases de datos"
	curso2.Slug = "BDD"
	curso2.Skills = []string{"Backend", "servidores"}

	otro := new(OtroCurso)
	otro.Name = "Nombre"

	carrera := new(Career)

	carrera.courses = []Course{curso}

	fmt.Println(curso)
	fmt.Println(curso2)
	fmt.Println(carrera)

	curso.Suscribe("Daniel")
}
