package main

func main (){
	salary:= 10000
	e:= myError{}
	if salary < 10000{

	}
}

type myError struct{}

func (e myError) Error() string {
	return "Error: el salario es menor a 10.000"
}

