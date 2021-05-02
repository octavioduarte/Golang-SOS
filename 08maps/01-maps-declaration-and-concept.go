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

	fmt.Println(members)
	fmt.Println(members["Janis Joplin"])
	fmt.Println(members["Jim Morrison"])

	// Adicionando novos valores

	members["Amy Winehouse"] = 2011

	fmt.Println(members)
	fmt.Println(members["Amy Winehouse"])
}
