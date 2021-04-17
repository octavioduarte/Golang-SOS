package main

import "fmt"

func main() {

	userProfile := "admin"

	switch {
	case userProfile == "root":
		fmt.Println("Usuário é root")
	case userProfile == "admin":
		fmt.Println("Usuário é admin")
	case userProfile == "guest":
		fmt.Println("Usuário é guest")
	}
}
