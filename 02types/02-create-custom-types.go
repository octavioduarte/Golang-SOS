package main

import "fmt"

type myCustomString string
type myCustomInt int

var (
	var_str myCustomString = "Woodstock"
	var_int myCustomInt    = 1969
)

func main() {
	fmt.Printf("Type of %v is %T\n", var_str, var_str)
	fmt.Printf("Type of %v is %T\n", var_int, var_int)
}
