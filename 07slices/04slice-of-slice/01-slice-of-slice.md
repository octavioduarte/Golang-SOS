> ## Slice of Slice 

Ja vimos que **um `Slice` tem o propósito de armazenar um conjunto de dados pertencentes a um contexto**, mas só vimos até agora ele armazenando `tipos primitivos`, porém é comum usarmos `Slices` com `Slices` internos, a sintaxe para isso em `Go` é a seguinte : 


```go
package main

import "fmt"

func main() {
    sliceOfSlice := [][]string{
        {"Go", "Node"},
        {"Python", ".Net"}
    }

	fmt.Println(sliceOfSlice)
}

// [[Go Node] [Python .Net]]
```

Não há um limite para os níveis internos de `Slices` dentro de outras `Slices`: 

```go
package main

import "fmt"

func sliceWtf() {
	sliceWtf := [][][][]string{
		{
			{
				{
					"Hello",
				},
			},
		},
		{
			{
				{
					"World !",
				},
			},
		},
	}

	fmt.Println(sliceWtf[0][0][0][0], sliceWtf[1][0][0][0])
}

// Hello World !

```