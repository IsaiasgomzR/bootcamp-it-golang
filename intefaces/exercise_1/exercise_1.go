package main

import "fmt"


type Product interface{
	Precio() float64
}
type small struct{
	precio float64
}

type medium struct{
	precio float64
}

type big struct {
	precio float64
}

func (s small) Precio() float64{
	return s.precio
}

func (m medium) Precio() float64 {
	return m.precio * 1.03
}

func (b big) Precio() float64{
	return b.precio * 1.06 + 2500
}

func Factory(size string, price float64) Product{
	switch size {
	case "small":
		return small{precio: price}
	case "medium":
		return medium{precio: price}
	case "big" :
		return big{precio: price}
	default: 
	return nil
	}

}

func main(){
	product := Factory("big", 100)
	precioTotal := product.Precio()
	fmt.Println(precioTotal)
}