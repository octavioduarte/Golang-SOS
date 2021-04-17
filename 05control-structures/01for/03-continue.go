package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		if x == 8 {
			continue
		}
		fmt.Println("Valor de x é", x)
	}
}
