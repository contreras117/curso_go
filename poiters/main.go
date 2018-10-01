package main

import (
	"fmt"
)

func main() {
	a := 234
	//Manera de declarar un variable de tipo apuntador
	//var b *int

	//El '&' se antepone a una variable para referirse a la direccion en memoria de esta.
	b := &a

	fmt.Println("Valores de las variables sin modificar")
	fmt.Println(a, *b)

	fmt.Println("Direcciones de las variables sin modificar")
	fmt.Println(&a, b)

	modifyPointer(b)
	fmt.Println("Valores de las variables modificadas")
	fmt.Println(a, *b)

}

func modifyPointer(c *int) {
	//El '*' se antepone a una variable de tipo apuntador para referirse al valor guardado en la direccion de memoria a la que apunta.
	*c = 129
}
