package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 5, 6, 7)

	fmt.Println(slice)
}
