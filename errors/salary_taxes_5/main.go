package main

import (
	"fmt"
	"errors"
)

var error1 = errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")

func monthSalary(hours int, value int) (float64, error){
	if hours < 80{
		return 0.0, fmt.Errorf("Error: %w", error1)
	}else{
	salary := hours * value
	if salary >= 150000 {
		return float64(salary) - (float64(salary)-  float64(salary) * 0.10), nil
	}else{
		return float64(salary), nil
	}
	}
}

func main(){
	var salary float64
	var err error
	salary, err = monthSalary(90,150) 
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("salary:" , salary)
	}
}
