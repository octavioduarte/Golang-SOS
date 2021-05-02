package main

import "fmt"

func main() {
	anonymousStruct := struct {
		age  int
		name string
	}{
		age:  27,
		name: "Janis",
	}

	fmt.Println(anonymousStruct)
}
