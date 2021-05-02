> ## Wft Slice 🤪 🤪 🤪

Existe uma particularidade com `slices` que pode complicar nossa vida quando programamos em `Go` : 



```go
package main

import "fmt"

func main() {
	slice1 := []int{2, 4, 6, 8, 10}

	fmt.Println("slice1 (before append of slice2)", slice1)

	slice2 := append(slice1[:2], slice1[3:]...)

	fmt.Println("slice1 (after append of slice2)", slice1)

	fmt.Println("slice2 = ", slice2)
}

// slice1 (before append of slice2) [2 4 6 8 10]
// slice1 (after append of slice2) [2 4 8 10 10]
// slice2 =  [2 4 8 10]
```

Se você reparar a saída, depois de realizarmos o `append` no `slice2`, você irá perceber que o `slice1` teve seu valor alterado 🤔 . 

Isso acontece pois quando o compilador verifica que `slice2` recebe uma parcial dos dados de `slice1` ele vai manipular os valores de `slice1` de forma que satisfaça o novo valor desejado para `slice2` . 

E em seguida "pega" os valores solicitados e atribui .

Isso pode causar grandes problemas no nosso código, pois podemos não ter a intenção de manipular os valores do `slice`, neste caso é recomendado seguir tratativas diferentes : 


Entre elas : 

Usando uma estrutura de repetição com uma condicional que insere apenas os valores pertinentes: 

> Range com Switch
```go
func usingRange() {
	slice1 := []int{2, 4, 6, 8, 10}
	slice2 := []int{}

	for _, value := range slice1 {
		switch value {
		case 2:
			fallthrough
		case 4:
			fallthrough
		case 8:
			fallthrough
		case 10:
			slice2 = append(slice2, value)
		}
	}

	fmt.Println("-----------------------------------------")
	fmt.Println("(usingRange) slice1 after range =", slice1)
	fmt.Println("(usingRange) slice2 =", slice2)
}

// (usingRange) slice1 after range = [2 4 6 8 10]
// (usingRange) slice2 = [2 4 8 10]
```

Se for possível (de acordo com o contexto do arquivo), podemos também manipular o próprio `Slice`, e teremos total controle do valor sobre ele : 

```go
func handleMainSlice() {
	slice1 := []int{2, 4, 6, 8, 10}

	slice1 = append(slice1[:2], slice1[3:]...)

	fmt.Println("-----------------------------------------")
	fmt.Println("(handleMainSlice) slice1 =", slice1)
}
// handleMainSlice) slice1 = [2 4 8 10]
```