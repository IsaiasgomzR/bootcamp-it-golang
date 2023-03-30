package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Iniciando la ejecucuion del programa")
	file, err := os.Open("customers.txt")

	if err != nil {
		panic("el archivo indicado no fue encontrado")
	}else{
		fmt.Println(&file)
	}

	defer func ()  {
		file.Close()
		fmt.Println("Ejecucion finalizada")
	}()
}