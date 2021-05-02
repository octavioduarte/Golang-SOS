package main

import "fmt"

type band struct {
	active  bool
	name    string
	country string
	year    int
}

func main() {
	acDc := band{
		active:  true,
		name:    "AC/DC",
		country: "Australia",
		year:    1973,
	}

	blackSabbath := band{false, "Black Sabbath", "England", 1968}

	fmt.Println(acDc)
	fmt.Println(blackSabbath)
}
