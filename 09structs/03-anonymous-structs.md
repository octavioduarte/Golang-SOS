> ## Anonymous Structs


`Structs anônimos` são `structs` onde a tipagem e a variável ficam vinculadas a uma única informação :



```go
package main

import "fmt"

func main() {
	anonymousStruct := struct {
		age  int
		name string
	}{
		age:  27,
		name: "Janis",
	}

	fmt.Println(anonymousStruct)
}



// {27 Janis}
```

Podemos omitir a chave (desde que passemos os valores na mesma sequência) : 


```go
package main

import "fmt"

func main() {
	anonymousStruct := struct {
		age  int
		name string
	}{
		27,
		"Janis",
	}

	fmt.Println(anonymousStruct)
}

// {27 Janis}
```