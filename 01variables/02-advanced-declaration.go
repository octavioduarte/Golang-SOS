package main

import "fmt"

var name, job, city string

var (
	item     string
	price    float64
	quantity int
)

func main() {
	message := "Hello"
	fmt.Printf("Valor de name: %v\r\n", name)
	fmt.Printf("Valor de job: %v\r\n", job)
	fmt.Printf("Valor de city: %v\r\n", city)
	fmt.Printf("Valor de item: %v\r\n", item)
	fmt.Printf("Valor de price: %v\r\n", price)
	fmt.Printf("Valor de quantity: %v\r\n", quantity)
	fmt.Printf("Valor de teste:  %v\r\n", message)
}
