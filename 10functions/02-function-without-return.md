> ## Functions without return


Em **todos os arquivos .go** declaramos uma **função sem retorno**  **neste repositório**, foi a função `main` , sua diferença básica em questões de sintaxe é que não informamos o **tipo de retorno** ao declarala-la e com isso o compilador ja entende que aquela função não irá retornar valor algum : 

```go
package main

import "fmt"

func main() {
	fmt.Println("I'm a function without return !!! ")
}

// I'm a function without return !!!
```