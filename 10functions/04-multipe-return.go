package main

import "fmt"


func main() {
	result, count := calc(1, 2, 3, 5, 7, 11, 13)

	fmt.Println(result, count)
}


func calc(numbers ...int) (int, int) {

	// numbers = [1 2 3 5 7 11 13]
	
	result := 0
	for _, v := range numbers {
		result += v
	}

	return result, len(numbers)
}