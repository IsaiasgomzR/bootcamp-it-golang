package main

import "fmt"


type Alumno struct{
	Name string
	LastName string
	DNI int
	BirthDate string
}
func (a Alumno) showDetail(){
	fmt.Printf("Nombre: %s\n", a.Name)
    fmt.Printf("Apellido: %s\n", a.LastName)
    fmt.Printf("DNI: %s\n", a.DNI)
    fmt.Printf("Fecha: %s\n", a.BirthDate)
}
func main(){
	alumno := Alumno{
		Name: "Isaias",
		LastName: "Gomez",
		DNI: 987654321,
		BirthDate: "18/06/****",
	}
	 alumno.showDetail()
	
}