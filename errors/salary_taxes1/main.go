package main

import "fmt"


func main(){
	var salary int = 140000
	err:= myError{}
	if salary < 150000 {
		fmt.Println(err.Error())
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}


type myError struct{}


func (e *myError) Error() string {
  return "Error: el salario ingresado no alcanza el mínimo imponible"
}