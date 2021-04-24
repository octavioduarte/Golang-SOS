# Arrays

> ## Sobre 

[Arrays](https://pt.wikipedia.org/wiki/Arranjo_(computa%C3%A7%C3%A3o)) em `Go` são poucos flexiveis, isso ja fica claro quando vemos que **seu tamanho não é dinâmico**, ou seja, **o tamanho passado na declaração não é mutavel** 

É assim que declaramos `Arrays` em `Go` : 

```go
package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
}

// [1 2 3 4 5]


/* 
  Na prática a sintaxe do exemplo acima se resume a : 

  variável x recebe um array de 5 elementos do tipo inteiro com os valores de 1 até 5
*/
```

> Podemos omitir os valores na declaração do array e neste caso o compilador atribuirá o valor 
`default` da tipagem 

```go 
package main 

import "fmt"

func main() {
    var x [5]int
    fmt.Println(x)
}

// [0 0 0 0 0 ]


// O mesmo exemplos com string ficaria da seguinte forma :  

package main 

import "fmt"

func main() {
    var x [5]string
    fmt.Println(x)
}

// [    ] Array com 5 strings vazias

```

>Se na atribuiçãod dos valores de um `Array` informarmos um número superior ao passado na declaração, teremos um erro 

```go
package main 

import "fmt"

func main() {
    var x [5]int
    x[0] = 1
    x[1] = 2
    x[2] = 3
    x[3] = 4
    x[4] = 5
    x[5] = 6

    fmt.Println(x)

    // invalid array index 5 (out of bounds for 5-element array)
}


// Ou

package main 

import "fmt"

func main() {
    x := [5]int{1, 2, 3, 4, 5, 6}
    fmt.Println(x)
}

// array index 5 out of bounds [0:5]
```