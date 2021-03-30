# Constantes

> ## Declaração 


- Em Go não há muita diferença entre a declaração e atribuição em constantes e variáveis, a sintaxe segue
praticamente o mesmo principio 


```code
package main

import "fmt"

const name string = "Janis Joplin"
const good_singer bool = true
const age int = 27

func main() {
	fmt.Println(name)
	fmt.Println(good_singer)
	fmt.Println(age)
}

# Output

// Janis Joplin
// true
// 27
```


> ## Notes

- Go permite a criação de constantes sem informar a tipagem . 

- O código abaixo funciona normalmente :


```code 
package main

import "fmt"

const name string = "Janis Joplin"
const good_singer bool = true
const age int = 27

func main() {
	fmt.Println(name)
	fmt.Println(good_singer)
	fmt.Println(age)

	declarationConstantWithoutType()
}

func declarationConstantWithoutType() {
	const const_without_type = "Hello"
	fmt.Println(const_without_type)
}

# Output

// Janis Joplin
// true
// 27
// Hello
```