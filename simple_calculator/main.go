package main

import "fmt"

const msgAskNumbers = "Ingresa los operandos separados por espacio: "

func main() {
	a, b := getNumbers()
	sum, res, mul, div := getOperations(a, b)
	fmt.Printf("Suma: %d | Resta: %d | Multiplicacion: %d | Division: %f", sum, res, mul, div)

}

func getNumbers() (int16, int16) {
	//para declarar enteros sin contar el signo se declara como uint
	//Por default los datos tipo Int son de 32 bits
	var a int16
	var b int16
	fmt.Print(msgAskNumbers)
	fmt.Scanf("%d %d", &a, &b)
	return a, b
}

func getOperations(a int16, b int16) (int16, int16, int16, float32) {
	sum, res, mul, div := (a + b), (a - b), (a * b), float32(a)/float32(b)
	return sum, res, mul, div
}
