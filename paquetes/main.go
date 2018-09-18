package main

import (
	"fmt"

	//Otra manera de hacer el import es con un punto antes del nombre. Usando esta forma no es
	//necesario escribir el nombre del paquete al invocar una funcion. Esta forma de import no es recomendada
	//y el linter te lo indica.
	//."github.com/contreras117/gocurso/paquetes/name"
	"github.com/contreras117/gocurso/paquetes/name"
)

func main() {
	nombre, apellido := name.GetName()
	fmt.Printf("Hola %s, %s", apellido, nombre)
}
