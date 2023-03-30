package main

import (
	"errors"
	"fmt"
)

type myError struct{
	err2 error
}
var error1 = errors.New("el mínimo disponible es de 150000 y el salario ingresado es de:")

func main (){

	var salary float64 = 140000 
	my_error := myError{err2: fmt.Errorf("Error: %w %f", error1, salary)}

	coincidence := false

	var err error

	if salary <= 150000{
		err = my_error.err2
	}

	coincidence = errors.Is(error1, errors.Unwrap(err))

	if coincidence{
		fmt.Println(err)
	}else{
		fmt.Println("TRUE")
	}
}
