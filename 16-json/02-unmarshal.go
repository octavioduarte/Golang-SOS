package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Name            string
	Category        []string
	Movie_theater   bool
	Sinopse         string
	Release_date_of string
}

func main() {
	jsonSliceByte := []byte(`{
		"Name":"Lord Of Rings Two Towers",
		"Category":[
		   "Action",
		   "Ação"
		],
		"Movie_theater":false,
		"Sinopse":"Sam is the real hero and Frodo gets in the way",
		"Release_date_of":"2002-12-27"
	 }`)

	var myStruct movie

	err := json.Unmarshal(jsonSliceByte, &myStruct)

	if err != nil {
		fmt.Println("Error on unmarshal", err)
	}

	fmt.Println(myStruct.Name)
}
