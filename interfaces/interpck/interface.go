//No es necesario que el package tengo el mismo nombre que el directorio en el que esta, pero es recomnndable.
package interfaces

//Platzi es una interface para usar metodos comunes de las structs Course y Career
type Platzi interface {
	Suscribe(name string)
}
