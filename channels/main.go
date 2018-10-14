//Los channels son unstructuras especiales usadas para enviar informacion entre diferentes hilos o rutinas.
//Al declararse se especifica el tipo de dato que manejaran, ya sean datos atomicos o complejos.
package main

import (
	"fmt"
)

//Un canal puede ser bidireccional o de un solo sentido, dependiende de la manera en
// en que se inicializa:
// chan tipoDato       Es un canal bidireccional
// chan<- tipoDato     Solo puede utilizar para enviar datos
// <-chan tipoDato     Solo puede utilizar para recibir datos

func main() {
	//channel simple que usa datos de tipo string.
	ch1 := make(chan string)
	go func() {
		//Con la funcion close no haceguramos que el channel ya no pueda recibir mas valores por seguridad en caso de un que otra
		//rutina en el futuro intente modificar los valores de ese channel.
		defer close(ch1)
		ch1 <- "Esto es un channel"
	}()

	fmt.Println(<-ch1)

	//channel simple que usa datos de tipo int. En este caso se le metieron multiples valores y para ir tomandolos se itero dentro
	//del channel con for y range. Usando la palabra range se tomaran los valores del channel hasta que se reciba la senal de que el channel ya esta cerrado.
	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := 1; i < 11; i++ {
			//Cuando se envia un dato al channel, el hilo que lo envia queda bloqueado hasta que el valor sea sacado del channel. En este caso el valor es tomado en el hilo main, dentro del for range (linea 33)
			//Lo anterior aplica mientras el bufer no este lleno (si no se declara un buffer, este tiene capacidad de 0)
			ch2 <- i
			fmt.Println("Se metio un dato al channel ch2")
		}
	}()
	for n := range ch2 {
		fmt.Printf("El valor %d fue tomado del channel ch2\n", n)
	}

	//De esta manera se declara un channel con buffer, es decir con un limite de valores que puede recibir. Una vez que se alcanza este limite
	//el channel no puede recibir mas valores hasta que los se tome algunos de los valores previos con <-
	ch3 := make(chan int, 3)
	ch3 <- 23
	ch3 <- 54
	ch3 <- 92
	fmt.Printf("El valor %d fue tomado del channel con buffer ch3\n", <-ch3)
	fmt.Printf("El valor %d fue tomado del channel con buffer ch3\n", <-ch3)
	fmt.Printf("El valor %d fue tomado del channel con buffer ch3\n", <-ch3)
	ch3 <- 81
	fmt.Printf("El valor %d fue tomado del channel con buffer ch3\n", <-ch3)

	ch4 := make(chan int, 3)
	f3 := func(ch chan int) {
		defer close(ch)
		ch <- 3
		ch <- 18
		ch <- 44
		ch <- 55
		ch <- 74
	}

	go f3(ch4)
	for n := range ch4 {
		fmt.Println(n)
	}

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(s[:len(s)/2], s[len(s)/2:])
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
