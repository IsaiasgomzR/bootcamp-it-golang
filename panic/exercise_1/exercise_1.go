package main


import (
	"fmt"
	"os"
)

func main(){

	file, err := os.Open("customers.txt")
	if err != nil{
		panic("the file not found or file be wrong")
	}else{
		fmt.Println(&file)
	}

	defer func ()  {
		file.Close()
		fmt.Println("Ejecucion finalizada")
	}()
	
}