package main

import "fmt"

type logs interface {
	say()
}

type human struct {
	name   string
	pharse string
}

func (h human) say() {
	fmt.Println("The human", h.name, "is speakin", h.pharse)
}

type dog struct {
	color  string
	pharse string
}

func (d dog) say() {
	fmt.Println("The dog is speakin", d.pharse)
}

func handleLogs(l logs) {
	l.say()
}

func main() {
	vinnyAppice := human{name: "Viny Appice", pharse: "Thanks Ronnie !!!"}
	bob := dog{color: "black", pharse: "AU AU AU"}

	handleLogs(vinnyAppice)
	handleLogs(bob)
}
