package main

import "fmt"

const (
	C1 = (iota + 1) * 5
	C2 = (iota + 2) * 10
	C3 = (iota + 3) * 20
	C4 = (iota + 4) * 40
)

func main() {
	fmt.Println(C1, C2, C3, C4)
}
