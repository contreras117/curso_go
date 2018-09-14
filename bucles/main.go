package main

import (
	"fmt"
)

func main() {
	forTest()
}

//En go la unica manera de hacer bucles es con For. No existe "While" ni "Do While".
func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("[FPR] ", i)
	}

	//Ciclo For esquivalente a un While. Solo declara la condicion, sin la asignasion de variable ni el autoincremento
	c := 100
	for c > 0 {
		fmt.Println("[FOR] solo con una condicion ", c)
		c -= 10
	}

	n := 1000
	for {
		n--
		if n == 0 {
			fmt.Println("Fin del ciclo 'Infinito'")
			break
		}
	}
}
