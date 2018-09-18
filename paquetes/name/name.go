package name

import "fmt"

//Para importar fuciones desde paquetes externos, su nombre debe iniciar en mayuscula para indicar que es publica.
//go exige que las funciones publicas deben ser documentadas.

//GetName Obtiene y devuelve el nombre y apellido capturados por el usuario.
func GetName() (string, string) {
	name := "Nombre"
	lastname := "Apellido"
	fmt.Print("Ingresa tu nombre y apellido paterno: ")
	fmt.Scanf("%s %s", &name, &lastname)
	return name, lastname
}
