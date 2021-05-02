> ## Deletando dados de uma slice

Ja deu pra notar que a equipe responsável pela criação do `Go` não curtia muito a ideia de diversas palavras reservadas vinculadas a uma linguagem . 

E isso vai ficar ainda mais evidente quando verificamos que **para manipulação de um `slice` o método append resolve boa parte dos casos**, a pricipio seu objetivo é mais voltado para criar novos indices em um `Array`, mas **também podemos usar esse mesmo méotod para `remover` um índice de um `Array`.**  



```go
package main

import "fmt"

func main() {
	queenMembers1991 := []string{"Brian", "John", "Freddie", "Roger"}

	queenMembers1991 = append(queenMembers1991[:2], queenMembers1991[len(queenMembers1991)-1:]...)

	fmt.Println(queenMembers1991)
}


// [Brian John Roger]
``` 

> OBS: 

Na prática o código acima irá recortar 2 pedaços do nosso `slice` ficando assim : 

```go
queenMembers1991[:2] // ==> [Brian, John]
queenMembers1991[len(queenMembers1991)-1:]... // ==> Roger

// O método len retorna um nº inteiro com o tamanho do Slice
```