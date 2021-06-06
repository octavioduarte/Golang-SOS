package main

import "fmt"

func main() {

	x := 10
	addressMemoryOfX := &x
	contentOfAddressOfMemoryOfX := *addressMemoryOfX

	fmt.Println("The content of", addressMemoryOfX, "is", contentOfAddressOfMemoryOfX)
}
