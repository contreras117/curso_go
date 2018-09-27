//No es necesario que el package tengo el mismo nombre que el directorio en el que esta, pero es recomnndable.
package interfaces

import (
	"fmt"
)

//Course - La palabra reservada struct no permite crear struct's, lo mas parecidos a clases dentro del leguaje GO.
type Course struct {
	Name   string
	Slug   string
	Skills []string
}

//Para crear funciones dentro que trabajen con una instancia de un struct (metodos), se crean como funciones colocando el struct que los va a utilizar antes del nombre de la funcion.

//Suscribe imprime en pantalla un mensaje de que el usuario proprorcionado se ha inscrito al curso que llamo al metodo.
func (c Course) Suscribe(name string) {
	fmt.Printf("El usuario %s se ha inscrito al curso %s. Felicidades!!!\n", name, c.Name)
}
