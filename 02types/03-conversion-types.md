# Tipos

> ## Conversão

- Go possibilita conversão de tipagem usando a seguinte sintaxe 

```go
    // TipagemDesejada(ValorParaConversão)
```


```go 
package main

import "fmt"

type customInt int

func main() {

	var x customInt = 4

	fmt.Printf("O tipo inicial de %v é %T\n", x, x)

	fmt.Println("Conversão realizada")

	fmt.Printf("Agora o tipo de %v é %T\n", x, int(x))

}

// O tipo inicial de 4 é main.customInt
// Conversão realizada
// Agora o tipo de 4 é int

```

