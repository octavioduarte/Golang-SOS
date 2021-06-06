package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name            string
	Category        []string
	Movie_theater   bool
	Sinopse         string
	Release_date_of string
}

func main() {
	lord_of_rings_two_towers := Movie{
		Name:            "Lord Of Rings Two Towers",
		Category:        []string{"Action", "Ação"},
		Movie_theater:   false,
		Sinopse:         "Sam is the real hero and Frodo gets in the way",
		Release_date_of: "2002-12-27",
	}

	mj, err := json.Marshal(lord_of_rings_two_towers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Movie JSON =", string(mj))
}
