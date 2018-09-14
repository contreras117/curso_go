package main

import (
	"fmt"
)

func main() {
	//En go las cedenas son manejadas en unicode (UTF8) de forma nativa.
	uniCodeString := getUnicode()
	uniCodeChar := getChar(uniCodeString, 2)

	fmt.Println("Cadena con UTF8: ", uniCodeString)
	fmt.Println(uniCodeChar)
	fmt.Println(string(uniCodeChar))
	fmt.Println("Cantidad de letras en la cadena: ", len(uniCodeString))

}

func getUnicode() string {
	return "カンジス"
}

func getChar(text string, position int8) byte {
	return text[position]
}
