> ## Maps 🔑💰 (Key and Value)


`Maps` é a forma como o `Go` gerencia a estrutura  `Chave Valor`, na prática ele é composto com um conjunto de dados como o `Slice`, mas tem uma estrutura um pouco mais objetiva e isso o torna mais eficiente :  



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

	fmt.Println(members)

    // Para acessar o valor basta informar o nome do map e a chave desejada
	fmt.Println(members["Janis Joplin"])
	fmt.Println(members["Jim Morrison"])
}

// map[Janis Joplin:1970 Jim Morrison:1971 Jimi Hendrix:1970 Kurt Cobain:1994 Robert Johnson:1938]
// 1970
// 1971
```