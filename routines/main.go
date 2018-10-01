//Las gouritines es la manera en que Go maneja los hilos o multiprocesos.

package main

import (
	"fmt"
	"time"
)

func main() {
	//Al usar "go" antes de la llamada de una funcion, se indica que esta funcion se ejecutara de manera paralela en otro hilo.
	go forGo(400)
	go forGo(240)
	//Si el hilo main termina, la ejecucion del programa se termina,
	//es por esto que se usa el sleep para que los hilos de las gorutines alcancen a terminar.
	time.Sleep(3 * time.Second)

}

func helloGo(n int) {
	fmt.Println("Esto es un print dentro de una goroutine #", n)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}
