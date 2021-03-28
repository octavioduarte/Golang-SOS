package main

import "fmt"

var (
	PublicVar, privateVar string
)

func main() {
	PublicVar = "Sou uma variável pública,  ( primeiro caratere é maiusculo )"
	privateVar = "Sou uma variável privada, ( primeiro caracatere é minúsculo )"
	fmt.Println("Variavel pública: ", PublicVar)
	fmt.Println("Variável privada: ", privateVar)
}
