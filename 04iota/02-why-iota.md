# Iota

> ## Quando usar 


- **Inteiros exigem menos armazenamento do que textos**, por tanto sempre que possível é ideal que o usemos no nosso código partindo desse principio é natural imaginar que nossas aplicações tenham armazenadas um conjunto de dados numéricos com o objetivo de identificar algo


- Vamos imaginar por exemplo que precisemos armazenar os níveis de acesso de usuários :


```go 
package main

import "fmt"

const (
	SIMPLE = iota
	ADMIN  = iota
	ROOT   = iota
)

func main() {
	fmt.Println(SIMPLE, ADMIN, ROOT)
}

// 0 1 2
```

