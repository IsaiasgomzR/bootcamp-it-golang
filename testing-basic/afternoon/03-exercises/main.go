package main

import (
	"fmt"
	"errors"
)

func main(){
	minutes := 20
	category:= "A"
	result, err := salaryFunction(category, minutesToHour(minutes))
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}

func salaryFunction(category string, hours float64 ) (total float64, err error){
	switch category {
	case  "A":
		total = (hours *3000) * 1.50
	case "B":
		total = (hours * 1500) * 1.20
	case "C":
		total = hours * 1000
	default: 
		errors.New("category no exists")
	}
	return
}

func minutesToHour( minutes int) (hours float64){
	const hour = 60
	return float64(minutes) / hour
}