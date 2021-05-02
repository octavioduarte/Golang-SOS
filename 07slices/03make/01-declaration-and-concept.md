> ## Make 

Um ponto importante sobre `Slices` é que **eles são muito mais convencionais, mas tem um custo computacional**,
o `Make` serve para otimizar esse custo que o `Slice` demanda . 

Na prática, informamos de forma prévia o `capacity` máximo que aquele `Slice` irá demandar, ou seja, o compilador sabe de forma antecipada o nível de processamento necessário para gerenciar o `Slice`.



```go
package main

import "fmt"

func main() {
	mySlice := make([]int, 5/*length*/, 10/*capacity*/)

	fmt.Println(mySlice, len(mySlice), cap(mySlice))
}

// [0 0 0 0 0] 5 10
```

O código acima diz que o `mySlice` é um slice de inteiros cujo o `length é igual a 5` e o seu `capcity é igual a 10`, ou seja,  a capacidade máxima que o slice ocupa é 10, o que significa que posso incluir até 5 elementos que o processamento não é afetado, pois essa informação já foi determinada.


**Isso não significa que não posso ultrapassar `capacity` previamente informado**, mas não teria sentido determinar um valor máximo para otimizar o processamento e não respeita-lo.

> ## O Slice sem make

Certo, ja sabemos que o **`Make` serve para informamos ao compilador o `capacity máximo` que um `Slice` irá ocupar**, mas como funciona um `Slice` sem o `Make` ?  E  porquê ele otimiza tanto o processamentoe ? 

Na prática **quando criamos um `Slice` (sem o make), o compilador atribui como `capacity` o `tamanho` do `Slice`**


```go
package main

import "fmt"

func main() {
	mySliceWithoutMake := []string{"Janis Joplin"}

	fmt.Println(cap(mySliceWithoutMake))
}
// 1
```

O `Slice` mySliceWithoutMake tem como `capacity` 1, ou seja, o compilador tem como informação prévia que aquele dado irá demandar um processamento suficiente para armazenar uma informação . 

Caso seja necessário expandir o meu `Slice` como o código abaixo faz: 

```go
package main

import "fmt"

func main() {
	mySliceWithoutMake := []string{"Janis Joplin"}

	fmt.Println("Capacity of mySliceWithoutMake is", cap(mySliceWithoutMake))

	// Expandindo o Slice

	mySliceWithoutMake = append(mySliceWithoutMake, "Woodstock 1969")

	fmt.Println("The new capacity of mySliceWithoutMake is", cap(mySliceWithoutMake))
}


// Capacity of mySliceWithoutMake if 1
// The new capacity of mySliceWithoutMake is 2
```

Por debaixo dos panos após o `append` na variável `mySliceWithoutMake` fará com que: 

* O compilador se desfaça da variável
* Crie uma nova com o valor atualizado  
* Atribua um novo `capacity` atribuido a ela (no caso 2)

Com o `capacity` previamente informado teria apenas uma etapa : 

* Atribua um novo valor a variável  