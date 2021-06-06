> ## JSON

**``JSON (Javascript Object Notation)``** é um formato de dados muito similar a objetos (que em `Go` chamamos de `structs`), por ser um conceito de uma outra tecnologia, mas amplamente utilizado por outras linguagens, é necessário fazer uma conversão de : 

Struct ➡️  JSON 
<br/>
<br/>
JSON   ➡️   Struct

```go
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
    // Marshal é o método do pacote json que usamos pra converter Structs para JSON

 	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Movie JSON =", string(mj))
    // O retorno do método Marshal é um slice of bytes, por tanto é necessário converter o retorno para o string . 

	/*  
		{
			"Name":"Lord Of Rings Two Towers",
			"Category":[
				"Action",
				"Ação"
			],
			"Movie_theater":false,
			"Sinopse":"Sam is the real hero and Frodo gets in the way",
			"Release_date_of":"2002-12-27"
		}
	*/
}
```

# IMPORTANTE

Conforme vimos na [sessão de visibilidade de dados](../01variables/03-public-and-privates.md) : 

**Dados com a letra maiúscula são publicos para outros pacotes**

<br/>

**Dados com a letra minuscula são privados para outros pacotes**

O pacote [encoding/json](https://golang.org/pkg/encoding/json/) é um pacote externo, por tanto, caso **você deseje converter um struct para JSON, você deve declarar o nome e as propriedades desse JSON com letra maiscula**