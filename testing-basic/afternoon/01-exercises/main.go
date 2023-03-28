package main 

import "fmt"


func main(){
  salary:= 50001
  result := taxSalary(salary)
  fmt.Println(result)
}

func taxSalary(salary int) float64 {
	var tax float64
	switch  {
	case salary > 50000:
		tax = float64(salary) * 0.17
	case salary > 150000:
		tax = float64(salary) * 0.27
	}
	return tax
}