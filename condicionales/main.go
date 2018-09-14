package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guestNumber()
}

func guestNumber() {
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed).Intn(10) + 1
	userNumber := 0

	fmt.Print("Adivina el nunero: ")
	fmt.Scanf("%d", &userNumber)

	if userNumber == randomNumber {
		fmt.Println("Adivinaste!")
	} else {
		fmt.Println("Suerte para la proxima!")
	}

	fmt.Println("User number: ", userNumber)
	fmt.Println("Random number: ", randomNumber)

	//Manera de declarar una vairable en la sentencia If
	if otherNumber := 3; otherNumber == 3 {
		fmt.Println("Entro en el IF")
	}
}
