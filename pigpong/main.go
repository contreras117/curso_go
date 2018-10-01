package main

import (
	"fmt"
)

func main() {
	end := make(chan bool)
	go pingpong(end)
	//El hilo main se bloquea hasta que haya un valor que tomar en el canal end
	<-end

}

func ping(ball chan<- int, action chan<- string) {
	ball <- 1
	action <- "Player Ping"

}

//Usando <-chan se indica que el parametro es un canal solo para enviar, es decir, la funcion no puede sacar valores del canal, solo meterlos.
func pong(ball chan<- int, action chan<- string) {
	ball <- 2
	action <- "Player Pong"
}

//Usando chan<- se indica que el parametro es un canal de solo lectura, es decir, la funcion no puede meter valores a este canal, solo sacarlos.
func referee(action <-chan string, end chan<- bool) {
	defer close(end)
	for i := 0; i < 1000; i++ {
		fmt.Printf("Action: %s\n", <-action)
	}
	end <- true
}

func pingpong(end chan<- bool) {
	ball := make(chan int)
	action := make(chan string)
	go referee(action, end)
	go ping(ball, action)
	for {
		switch <-ball {
		case 1:
			go pong(ball, action)
		case 2:
			go ping(ball, action)
		}
	}

}
