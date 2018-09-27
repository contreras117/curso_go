//No es necesario que el package tengo el mismo nombre que el directorio en el que esta, pero es recomnndable.
package interfaces

import "fmt"

//Career es un struct que en sus valores internos tiene un arreglo de otro struct
type Career struct {
	Name    string
	Courses []Course
}

//Suscribe es un metodo que suscribe a un alumno a una carrera.
func (c Career) Suscribe(name string) {
	fmt.Printf("El usuario %s se ha inscrito a la carrera  de %s. Felicidades!!!\n", name, c.Name)
}
