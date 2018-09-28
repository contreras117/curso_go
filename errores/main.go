package main

import (
	"fmt"

	"github.com/contreras117/gocurso/errores/math"
)

func main() {
	//En go no existe las execpciones, y la manera de detectar errores es validando si el valor error es igual a nil y controlando el flujo del programa en base a esto.
	number, err := math.Sum(50, 40)

	if err != nil {
		//La palabra panic detiene el flujo del programa inmediatamente e imprime en pantalla el mensaje del error.
		panic(err)
	}
	fmt.Println(number)

	number, err = math.Sum(50, "40")

	//Esta es la manera de manejar errores en Go, si nil es diferente de nil, se maneja el error, sino, el flujo del programa sigue normalmente.
	if err != nil {
		panic(err)
	}
	fmt.Println(number)

}
