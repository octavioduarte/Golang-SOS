# Else

> ## Sobre 

- O `else` é um complemento a estrutura `if` que vai determinar o trecho de código a ser executado caso a condição seja falsa


Exemplo : 

<br>

```go
package main

import "fmt"

func main() {
	clubWorldTitles := 2

	if clubWorldTitles == 0 {
		fmt.Println("🌴")
	} else {
		fmt.Println("🦅")
	}
}

// 🦅
```

