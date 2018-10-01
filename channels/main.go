//Los channels son unstructuras especiales usadas para enviar informacion entre diferentes hilos o rutinas.
//Al declararse se especifica el tipo de dato que manejaran, ya sean datos atomicos o complejos.
package main

import (
	"fmt"
)

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
	//del channel con for y range, esta ultima es similar a foreach en Java.
	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := 1; i < 11; i++ {
			ch2 <- i
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
	f3 := func(ch *chan int) {
		defer close(*ch)
		*ch <- 3
		*ch <- 18
		*ch <- 44
		*ch <- 55
		*ch <- 74
	}

	go f3(&ch4)

	for n := range ch4 {
		fmt.Println(n)
	}
}
