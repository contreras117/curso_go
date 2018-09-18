package main

import "fmt"

func main() {
	switchTest()
}

func switchTest() {
	number := 0
	fmt.Print("[Switch] Ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("El numero es 1")
	default:
		fmt.Println("El numero no es 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}

	// A diferencia de otros lenguajes, en go no es necasario usar la palabra break al final de cada case.
	// Switch solo entrara en el primer casp que cumpla su condicion.
	switch {
	case number > 10:
		fmt.Println("El numero es mayor a 10")
	case number > 100:
		fmt.Println("El numero es mayor a 100")
	default:
		fmt.Println("El numero no es mayor a 10")
	}
}
