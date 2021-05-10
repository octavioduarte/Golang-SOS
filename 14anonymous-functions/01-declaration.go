package main

import "fmt"

func main() {

	woodstockYear := 1969

	func(year int) {
		fmt.Println("The best woodstock happen in", year)
	}(woodstockYear)
}
