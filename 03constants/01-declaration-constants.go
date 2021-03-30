package main

import "fmt"

const name string = "Janis Joplin"
const good_singer bool = true
const age int = 27

func main() {
	fmt.Println(name)
	fmt.Println(good_singer)
	fmt.Println(age)

	declarationConstantWithoutType()
}

func declarationConstantWithoutType() {
	const const_without_type = "Hello"
	fmt.Println(const_without_type)
}
