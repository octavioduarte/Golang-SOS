> ## Functions with multiple return

Quando o retorno de nossas função não é apenas um [`tipo primitivo`](https://github.com/octavioduarte/Golang-SOS/blob/master/02types/01-primitive-types.md) e desejamos retornar diversos valores, basta que adicionemos entre parenteses todas as tipagens desejadas


```go 
package main

import "fmt"


func main() {
	result, count := calc(1, 2, 3, 5, 7, 11, 13)

	fmt.Println(result, count)
}

                           // Significa que a função calc retorna 2 números inteiros
func calc(numbers ...int) (int, int) {

	// numbers = [1 2 3 5 7 11 13]
	
	result := 0
	for _, v := range numbers {
		result += v
	}

	return result, len(numbers)
}

// 42 7

```