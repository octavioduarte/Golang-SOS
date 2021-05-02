> ## Range with Maps


O método `range` tem uma sintaxe identica em `slices` e `maps`, a diferença é que com **`Maps`, os parâmetros são `Chave` e `Valor`**, pois em `Maps` não temos índices



```go
package main

import "fmt"

func main() {

	members := map[string]int{
		"Robert Johnson": 1938,
		"Janis Joplin":   1970,
		"Jimi Hendrix":   1970,
		"Jim Morrison":   1971,
		"Kurt Cobain":    1994,
	}

	for key, value := range members {
		fmt.Println("Key: ", key, "has value", value)
	}
}

// Key:  Robert Johnson has value 1938
// Key:  Janis Joplin has value 1970
// Key:  Jimi Hendrix has value 1970
// Key:  Jim Morrison has value 1971
// Key:  Kurt Cobain has value 1994
```