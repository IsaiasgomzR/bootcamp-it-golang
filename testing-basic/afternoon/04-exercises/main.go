package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main(){
	oper := operation(minimum)
	result := oper(1,2,3)
	fmt.Println(result)
}

func operation(op string) func(values ...int) float64{
	switch op{
	case maximum :
      return maxFunc
	case average :
		return averageFunc
	case minimum:
		return minFunc
	}
	return nil
}

func minFunc(values ...int) (minValue float64){
	maxValue := maxFunc(values...)
	for _, value := range values {
		if value < int(maxValue){
			maxValue = float64(value)
		}
	}
	return float64(maxValue)

}

func maxFunc( values ...int) (maxValue float64){
	for _, value := range values {
		if value > int(maxValue){
			maxValue = float64(value)
		}
	}
	return
}

func averageFunc ( values ...int) (total float64){
	var sum int
	for _, value := range values {
		sum += value
	}
	return float64((sum / len(values)))
}