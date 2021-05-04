package main

import "fmt"

func main() {
	message := printMessage()
	number := printNumber()

	fmt.Println(message)
	fmt.Println(number)
}

func printMessage() string {
	return "A simple funcion with Go !!! "
}

func printNumber() int {
	return 1969
}
