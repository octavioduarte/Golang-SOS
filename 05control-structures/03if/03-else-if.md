# Else

> ## Sobre 

- O `else if` é também  um complemento a estrutura `if` a diferença é que ao contrário do **`else` que executa de forma automática o trecho imposto a ele sem quaisquer análises**, ele vai analisar previamente uma determinada condição e só em caso `verdadeiro` ele irá executá-la  


Exemplo : 

<br>

```go
package main

import "fmt"

func main() {
	clubWorldTitles := 0

	if clubWorldTitles == 2 {
		fmt.Println("🦅")
	} else if clubWorldTitles == 0 {
		fmt.Println("🌴")
	}
}

// 🌴
```

