package main 

import (
	"fmt"
	"errors"
)
type Customer struct{
	DNI int
	Name string
	address string
	phone int
	Legago string
}

func (c *Customer) Save (){
	customersList = append(customersList, *c)
}

var customersList = []Customer{
	{
		Name: "Isaias",
		address: "fernon ST",
		phone: 987654321,
		Legago: "null",
		DNI: 123456789,
	},
}

func main() {
	fmt.Println("Iinico de la ejecucuion del program")

	newCustomer:= Customer{
		Name: "Iconic",
		address: "fernon ST",
		phone: 987654321,
		Legago: "null",
		DNI: 123456789,
	}

	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println("se detectaron varios errores en tiempo de ejcucuoin")
			fmt.Println("Recoverd:", err)
		}
		fmt.Println("Finalizacion de programa")
	}()

	_, err := validateCustomer(&newCustomer	)
	if err != nil{
		panic(err)
	}
	for _, item := range customersList {
		if item.DNI == newCustomer.DNI{
			panic("Error: el cliente ya existe")
		}
	}

	newCustomer.Save()
	fmt.Println(customersList)
}

func validateCustomer (c *Customer) (bool ,error){
	if c.Name == "" || 
	c.DNI == 0 || 
	c.address == "" ||
	 c.phone == 0 || 
	 c.Legago == "" {
		return false, errors.New("Error: todos los campos deben ser completos")
	 }
	 return true, nil
 }