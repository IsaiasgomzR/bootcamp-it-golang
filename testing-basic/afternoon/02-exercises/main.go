package main

import "fmt"

func main (){
	note := []int {1,2,3,4}
	fmt.Println( averageNotes(note...))

}

func averageNotes(notes ...int) (average float64){
	
	for _, value := range notes {
		average += float64(value)
	}
	return float64(int(average) / len(notes))
}