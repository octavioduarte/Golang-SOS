package main

import "fmt"

type dog struct {
	name  string
	age   int
	color string
}

func (d dog) printAge() {
	fmt.Println(d.name, "has", d.age, "years")
}

func (d dog) printColor() {
	fmt.Println("The color of", d.name, "is", d.color)
}

func (d dog) printBark() {
	fmt.Println(d.name, "is talking AU AU AU")
}

func main() {
	dog1 := dog{"Guga", 18, "White"}
	dog2 := dog{"Jadi", 20, "Black"}
	dog3 := dog{"Bob", 4, "Gray"}

	dog1.printAge()
	dog2.printBark()
	dog3.printColor()
}
