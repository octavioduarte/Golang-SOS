package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)

	fmt.Println(mySlice, len(mySlice), cap(mySlice))
}

func sliceWithoutMake() {
	mySliceWithoutMake := []string{"Janis Joplin"}

	fmt.Println("Capacity of mySliceWithoutMake is", cap(mySliceWithoutMake))

	// Expandindo o Slice

	mySliceWithoutMake = append(mySliceWithoutMake, "Woodstock 1969")

	fmt.Println("The new capacity of mySliceWithoutMake is", cap(mySliceWithoutMake))
}
