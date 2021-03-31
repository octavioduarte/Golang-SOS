package main

import "fmt"

const (
	SIMPLE = iota // 0
	ADMIN  = iota // 1
	ROOT   = iota // 2
)

func main() {
	fmt.Println(SIMPLE, ADMIN, ROOT)
}
