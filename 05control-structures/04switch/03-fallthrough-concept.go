package main

import "fmt"

func main() {
	statusPayment := "error - insufficient funds"

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
