package main

// Explicit typing
var name string = "Janis Joplin"
var good_singer bool = true
var age int = 27

// Without explicit typing
var city = "Port Arthur"

var document int   // Será atribuído   ==> 0
var country string // Será atribuído   ==> ""
var logged bool    // Será atribuído   ==> false

func main() {
	println(name)
	println(good_singer)
	println(age)
	println(city)
	println(document)
	println(country)
	println(logged)
}
