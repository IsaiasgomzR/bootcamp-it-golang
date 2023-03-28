package main

import "fmt"


var products = []Product {
	Product{
		ID: 1,
		Name: "Gomitas",
		Price: 15.0,
		Description: "good",
		Category: "Dulces",
	},
	Product{
		ID: 2,
		Name: "Gomitas",
		Price: 15.0,
		Description: "good",
		Category: "Dulces",
	},
}

type Product struct{
	ID int
	Name string
	Price float64
	Description string
	Category string
	
}


func (p *Product) save(){
 products =	append(products, *p)
}

func (p *Product) getAll() []Product {
	return products
}

func (p *Product) findById(ID int) (target Product){
	for _, item := range products {
		if item.ID == p.ID {
			target = item
		}
	}
	return
}
 
/////////////////////////////////


type Person struct{
	ID int
	Name string
	DateOfBirth string
}

type Employee struct{
	ID int
	Position string
	Person
}

func (e *Employee) printEmployee(){
	fmt.Println(e.DateOfBirth, e.Name , e.Position, e.ID)
}
 

func main(){
	p1 := Product{
		ID: 3,
		Name: "other",
		Price: 12.0,
		Category: "Other",
		Description: "Other",
	}
	p1.save()
	fmt.Println(p1.getAll())

	person1 := Person{
		ID: 1,
		Name: "Isaias",
		DateOfBirth: "06/18/****",
	}
	employee1 := Employee{
		ID: 1,
		Position: "Dev",
		Person: person1,
	}

	employee1.printEmployee()

}

