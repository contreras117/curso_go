package math

import "errors"

//Sum adds two values if both are numbers.
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("Error!!!, el valor A es un string")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("Error!!!, el valor B es un string")
	}
	return a.(int) + b.(int), nil
}

//Multiplication multiplicates two values if both are numbers.
func Multiplication(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("Error!!!, el valor A es un string")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("Error!!!, el valor B es un string")
	}
	return a.(int) * b.(int), nil
}
