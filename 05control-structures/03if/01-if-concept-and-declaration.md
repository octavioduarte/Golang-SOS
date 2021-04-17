# If

> ## Sobre 

 Existem basicamente duas estruturas de controle as de **repetição** e as de **condição**.

* **As estruturas de repetição (como o `for` que ja vimos) que vão repetir um determinado trecho de código até que uma determinada condição seja alcançada .**

<br>

* **As de condição (que é o caso do `if`) que vai analisar uma determiada a condição e seguir um "rumo" condicionado ao resultado dela .**


Exemplo : 

<br>

```go
package main

import "fmt"

func main() {
	clubWorldTitles := 0

	if clubWorldTitles == 0 {
		fmt.Println("🌴")
	}

	// 🌴
}
```

- É uma sintaxe relativamente simples, que se resume a : 

```code
	if (// condição) {
		// trecho a ser executado caso a condição seja verdadeira
	}
```
