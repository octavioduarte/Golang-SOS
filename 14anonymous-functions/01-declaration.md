> ## Funções Anônimas 🎩

`Funções anônimas` em `Go` são funções em que omitimos o nome e tem o propósito de **ser executada apenas uma vez, ou seja, seu uso é descartável** 


```go
package main

import "fmt"

func main() {

	woodstockYear := 1969

	func(year int) {
		fmt.Println("The best woodstock happen in", year)
	}(woodstockYear) // Isso invoca a função
}

// The best woodstock happen in 1969
```

Note que o parenteses após o bloco da função, isso a torna **auto invocável**, pois como não há um nome para invoca-la, sua sintaxe **exige** que ela seja chamada dessa forma quando lida pelo compilador