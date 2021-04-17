package main

import "fmt"

func main() {
	country := "Brasil"

	switch country {
	case "China":
		fmt.Println(country, "fica na Ásia")
	case "Brasil":
		fmt.Println(country, "fica na America do Sul")
	case "Russia":
		fmt.Println(country, "fica na Europa e Ásia")
	}
}
