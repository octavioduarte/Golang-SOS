# Constantes

> ## Constantes e tipagem 


- Verificamos na [documentação de declaração de constantes](https://github.com/octavioduarte/Golang-SOS/blob/master/03constants/01-declaration-constants.md) que 

**Go permite a criação de constantes sem atribuição de tipagem e isso também é possível com variáveis**,

**A diferença é que em constantes quando a tipagem não é está na declaração o compilador do Go não atribui** **tipagem nessa constante até que ela seja usada**

 Por exemplo se você declarar uma constante na linha `03` do seu arquivo sem atribuir uma tipagem, o compilador não saberá do que se trata aquela informação até que ela seja usada .


```go
package main

import "fmt"

const const_without_type = 10

var x float64 = const_without_type

func main() {
	fmt.Println(x)
}

// 10
```

- Estamos atribuindo a `variável x` que está tipada como `float64` um valor `int`, como foi dito isso funciona pq o compilador não "sabe" até o momento da atribuição qual era a tipagem de "const_without_type"


**O mesmo exemplo com variáveis, irá resultar em erro :**  

```go
package main

import "fmt"

var var_without_type = 10

var x float64 = var_without_type

func main() {
	fmt.Println(x)
}

// cannot use var_without_type (type int) as type float64 in assignment
```