package main

import "fmt"

func main() {
	slice1 := []int{2, 4, 6, 8, 10}

	fmt.Println("slice1 (before append of slice2)", slice1)

	slice2 := append(slice1[:2], slice1[3:]...)

	fmt.Println("slice1 (after append of slice2)", slice1)

	fmt.Println("slice2 = ", slice2)

	usingRange()
	handleMainSlice()
}

func usingRange() {
	slice1 := []int{2, 4, 6, 8, 10}
	slice2 := []int{}

	for _, value := range slice1 {
		switch value {
		case 2:
			fallthrough
		case 4:
			fallthrough
		case 8:
			fallthrough
		case 10:
			slice2 = append(slice2, value)
		}
	}

	fmt.Println("-----------------------------------------")
	fmt.Println("(usingRange) slice1 after range =", slice1)
	fmt.Println("(usingRange) slice2 =", slice2)
}

func handleMainSlice() {
	slice1 := []int{2, 4, 6, 8, 10}

	slice1 = append(slice1[:2], slice1[3:]...)

	fmt.Println("-----------------------------------------")
	fmt.Println("(handleMainSlice) slice1 =", slice1)
}
