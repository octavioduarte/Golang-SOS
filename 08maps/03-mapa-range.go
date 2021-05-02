package main

import "fmt"

func main() {

	members := map[string]int{
		"Robert Johnson": 1938,
		"Janis Joplin":   1970,
		"Jimi Hendrix":   1970,
		"Jim Morrison":   1971,
		"Kurt Cobain":    1994,
	}

	for key, value := range members {
		fmt.Println("Key: ", key, "has value", value)
	}
}
