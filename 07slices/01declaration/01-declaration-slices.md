# Slices

> ## Sobre 

[Slices](https://tour.golang.org/moretypes/7) em `Go` são uma forma mais flexível do que `Arrays` para criar dados agrupados, sua diferença básica em relação a `Arrays` é que o tamanho inicial de sua declaração é mutável, além disso **ao contrário da declaração de `Arrays`, com `Slices` não precisamos informar o tamanho inicial dos nossos dados .** 


```go
package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Array = ", myArray)
	fmt.Println("Slice = ", mySlice)
}

// Array =  [1 2 3 4 5]
// Slice =  [1 2 3 4 5]
```
