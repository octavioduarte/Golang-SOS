package main

import "fmt"

var (
	var_string string  = "Janis"
	var_int    int     = 27
	var_bool   bool    = true
	var_float  float64 = 10.5
)

func main() {
	fmt.Printf("Type of %v is %T\n", var_string, var_string)
	fmt.Printf("Type of %v is %T\n", var_int, var_int)
	fmt.Printf("Type of %v is %T\n", var_bool, var_bool)
	fmt.Printf("Type of %v is %T\n", var_float, var_float)
}
