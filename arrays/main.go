package main

import "fmt"

func main() {
	getArray()
	getSlice()

}

//La diferencia entre los Array y los slice es que los Array son de tamano fijo mientras que los
//Slice son de tamno dinamico.
//Al crear un slice este no tiene ninguna posicion dentro de el, por lo que no se pueden usar
//indices para asignar nuevos valores en mediante ellos. Para meter un nuevo valor a un Slice
// se usa la funcion append.
func getArray() {
	var arr1 [2]string
	arr2 := [3]int8{1, 2, 3}
	arr1[0] = "first"
	arr1[1] = "second"
	fmt.Println(arr1, arr2)
}

func getSlice() {
	var slice1 []string
	slice2 := []int8{1, 2, 3, 4, 5}
	slice2[3] = 9
	slice1 = append(slice1, "first", "second", "third")
	fmt.Println(slice1, slice2)
	fmt.Println(slice1[2], slice2[1])
}
