package main

import "fmt"

func main() {
	queenMembers1991 := []string{"Brian", "John", "Freddie", "Roger"}

	queenMembers1991 = append(queenMembers1991[:2], queenMembers1991[len(queenMembers1991)-1:]...)

	fmt.Println(queenMembers1991)
}
