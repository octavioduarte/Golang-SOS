> ## Structs


`Structs` é serve para agruparmos dados de diferentes tipos pertencentes a um contexto em comum.

Na prática criamos sua `"interface"`, ou seja **declaramos os campos e seus respectivos tipos, e as variáveis instanciadas vinculadas a esse `Struct` devem respeitar o que foi determinado.**



```go
package main

import "fmt"

type band struct {
	active  bool
	name    string
	country string
	year    int
}

func main() {
	acDc := band{
		active:  true,
		name:    "AC/DC",
		country: "Australia",
		year:    1973,
	}

    // Podemos omitir as chaves, porem devemos respeitar exatamente a ordem estabelecida no Struct
	blackSabbath := band{false, "Black Sabbath", "England", 1968}

	fmt.Println(acDc)
	fmt.Println(blackSabbath)
}


// {true AC/DC Australia 1973}
// {false Black Sabbath England 1968}
```