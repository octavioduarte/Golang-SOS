> ## Range 

O método [Range](https://tour.golang.org/moretypes/16) em `Go` é um método voltado para `Slices` que tem o propósito de "varrer" todo os dados agrupados . 




```go
package main

import "fmt"

func main() {
	members := []string{"Janis", "Hendrix", "Morrison", "Robert"}

	for index, value := range members {
		fmt.Println("Index", index, "value is", value)
	}
}

/*
    Index 0 value is Janis
    Index 1 value is Hendrix
    Index 2 value is Morrison
    Index 3 value is Robert
*/
```