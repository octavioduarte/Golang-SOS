package main

import "fmt"

func main() {
	members := map[string]int{
		"Robert Johnson": 1938,
		"Janis Joplin":   1970,
		"Jimi Hendrix":   1970,
		"Jim Morrison":   1971,
		"Kurt Cobain":    1994,
		"Fake Member":    0,
	}

	valueOzzy, okOzzy := members["Ozzy Osbourne"]
	fmt.Println(valueOzzy, okOzzy)

	valueFakeMember, okFakeMember := members["Fake Member"]
	fmt.Println(valueFakeMember, okFakeMember)
}
