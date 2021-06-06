> ## JSON - Unmarshal

Da mesma forma que é necessário fazer a conversão de um `Strunct` para `JSON` quando vamos nos comunicar com um servidor que aceita apenas este formato, devemos fazer também a conversão de `JSON` para `Strunct` quando um servidor `Go` é o receptor das informações, pois **`Go não manipula nativamente JSON`** . 

Este códidgo converte um `JSON` para `Strunct`

```go
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
    // É necessário primeiramente converter o JSON para um Slice of Bytes
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

    // Esta é a variável que vai receber os valores convertidos
	var myStruct movie
                                         // Devemos passar o ponteiro da variável
	err := json.Unmarshal(jsonSliceByte, &myStruct)

	if err != nil {
		fmt.Println("Error on unmarshal", err)
	}

	fmt.Println(myStruct.Name)

    // Lord Of Rings Two Towers
}
```