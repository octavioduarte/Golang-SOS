package main

import "fmt"

func main() {

	x := 10
	addressMemoryOfX := &x

	fmt.Println("The address memory of variable x is ", addressMemoryOfX)
}
