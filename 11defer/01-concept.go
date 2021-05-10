package main

import "fmt"

func main() {
	defer fmt.Println(execCount())
	fmt.Println("Janis Joplin passed away with: ")
}

func execCount() int {
	return 1970 - 1943
}
