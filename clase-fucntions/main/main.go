package main 

suma := "+"
resta := "-"
multipiclacion := "*"
division := "/"

func main (){
result := fucntionAritmetica(5,5,suma)
fmt.println(result)
}


func fucntionAritmetica(value1 float64 ,value2 float64, operator string) float64 {
	switch(operator){
	case suma :
		 return	value1 + value2 	
	case resta :
			return value1 - value2
	case multipiclacion :
			return value1 * value2
		
	case division :
			return value1 / value2
		
		 return 0
	}
}