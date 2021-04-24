package main

import "fmt"

func main() {
	members := []string{"Janis", "Hendrix", "Morrison", "Robert"}

	for index, value := range members {
		fmt.Println("Index", index, "value is", value)
	}
}
