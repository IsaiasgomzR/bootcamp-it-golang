package main

import (
	"fmt"
)

func main(){
	tickets.readFile("tickets.csv")

	total, err:= tickets.GetTotalTickets("Mexico")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(total)
	total , _ := tickets.GetMornings("madrugada")
	println(total)

	total,_ := tickets.GetMornings("manana")
	println(total)

	total,_ := tickets.GetMornings("tarde")
	println(total)

	total,_ := tickets.GetMornings("noche")
	println(total)

	totales, err := tickets.AverageDestiantion("Brazil", 10)
	if err != nil {
		println(err)
	}
	println(totales)
}