package main

import (
	"fmt"
	"strings"
)

func main() {
	cadenas()
}

func cadenas() {
	text := "Hello world, Hello Platzi, Hello go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ","))
}
