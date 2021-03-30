package main

import "fmt"

const const_without_type = 10

var x float64 = const_without_type

func main() {
	fmt.Println(x)
}
