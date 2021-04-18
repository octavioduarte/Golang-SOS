# Fallthrough


## Sobre 

O `fallthrough` é usado quando duas ou mais possibilidades de análise do meu `switch` devem seguir exatamente o mesmo trecho de código : 



```go
package main

import "fmt"

func main() {
	statusPayment := "error - insufficient funds"

/*
     Neste cenário é indiferente saber se o erro que ocorreu é uma insuficiência de saldo,
ou se o banco recusou o pagamento, o que me importa é notificar que um erro ocorreu com uma mensagem genérica

*/ 
	switch statusPayment {

	case "ok":
		fmt.Println("Payment worked")

	case "pending":
		fmt.Println("Payment is pending")

	case "error - insufficient funds":
		fallthrough

	case "error - denied by the bank":
		fallthrough

	case "error":
		fmt.Println("Payment failed")
	}
}


// Payment failed
```