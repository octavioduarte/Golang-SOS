# For

> ## Declaração 

- A **estrutura de repetição** `for` não foge muito do "padrão" da grande maioria das linguagens, sua composição se baseia em :

* `Inicializador` -->  Valor inicial do contador
* `Condição` -->       Condição de saída da estrutura 
* `Pós` -->            O que acontecerá para cada [loop](https://pt.wikipedia.org/wiki/Loop_(programa%C3%A7%C3%A3o)) da estrutura


```go
package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}



/* 
  0
  1
  2
  3
  4
  5
  6
  7
  8
  9
*/
```

**[Link util](https://gobyexample.com/for)**