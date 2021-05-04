> ## Basic functions

`Go` tem uma sintaxe para escrita de funções bem similar a outras **linguagens tipadas**, onde os principais componentes são : 

Uma **palavra reservada** para determinar que aquele código refere-se a uma `função`
O **nome** da `função`
O(s) **parâmetro(s)** com suas respectivas tipagens
O **tipo do retorno**


Exemplo : 

```go
package main

import "fmt"

func main() {
    message := printMessage()

    fmt.Println(message)
}


func printMessage() string {
	return "A simple funcion with Go !!! "
}

// A simple funcion with Go !!! 
```
Onde : 

| Nome             | Descrição                                                                      |                      
| ---------------- | ------------------------------------------------------------------------------- | 
| `func`           | **palavra reservada**                                                           |
| `printMessage`   | **nome**                                                                        |                     
| `()`             | **parâmetro(s)** (neste caso não há nenhum, mas ainda assim usamos o parenteses)       |                     
| `string`         | **tipo do retorno**                                                                              |                     


