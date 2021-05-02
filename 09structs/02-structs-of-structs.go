package main

import "fmt"

type singerInfo struct {
	born int
	dead bool
	name string
}

type band struct {
	active bool
	singerInfo
	name    string
	country string
	year    int
}

func main() {
	dio := band{
		active: false,
		singerInfo: singerInfo{
			name: "Ronnie James Dio",
			born: 1942,
			dead: true,
		},
		name:    "DIO",
		country: "EUA",
		year:    1982,
	}

	fmt.Println(dio)
}
