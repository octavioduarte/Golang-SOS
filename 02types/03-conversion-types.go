package main

import "fmt"

type customInt int

func main() {

	var x customInt = 4

	fmt.Printf("O tipo inicial de %v é %T\n", x, x)

	fmt.Println("Conversão realizada")

	fmt.Printf("Agora o tipo de %v é %T\n", x, int(x))

}
