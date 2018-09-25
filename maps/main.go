package main

import (
	"fmt"

	"github.com/contreras117/gocurso/maps/maps"
)

func main() {

	testMap := maps.GetMapMake()
	fmt.Println(testMap)
	testMap = maps.DeleteValue(testMap, "llave2")
	fmt.Println(testMap)
	fmt.Println(maps.GetEdad("Daniel"))

	//Si la clave no esta dentro del mapa, regresa 0 por default
	fmt.Println(maps.GetEdad("Marco"))

}
