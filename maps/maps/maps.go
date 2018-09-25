package maps

//GetMapMake regresa un mapa inicializado con la funcion make
func GetMapMake() map[string]int {
	//funcion para crear e inicializar un map. Si no se inicializa no se puede utilizar.
	testMap := make(map[string]int)
	testMap["llave1"] = 1
	testMap["llave2"] = 2
	testMap["llave3"] = 3
	return testMap
}

//GetEdad regresa la edad (valor) correspiendente dentro del mapa al nombre proporcionado (llave)
func GetEdad(nombre string) int {
	mapaEdad := map[string]int{
		"Daniel": 26,
		"Carlos": 28,
		"Andrea": 21,
	}

	return mapaEdad[nombre]
}

//DeleteValue regresa el mapa pasado como parametro borrando el valor de correspondiente a la llave pasada como parametro
func DeleteValue(mapa map[string]int, llave string) map[string]int {
	delete(mapa, llave)
	return mapa
}
