package main

import (
	"fmt"
	"errors"
)

func main (){
	salary:= 10000
	my_error := myError{}

	var err error
	coincidence := false
	if salary < 10000{
		err = my_error.Error()
	}

	coincidence = errors.Is(err, error1)

	if coincidence{
		fmt.Println(err)
	}else{
		fmt.Println("TRUE")
	}
}

type myError struct{}

var error1 = errors.New("el salario es menor a 1000")

func (e myError) Error() error {
	return fmt.Errorf("Error: %w",error1  )
}



