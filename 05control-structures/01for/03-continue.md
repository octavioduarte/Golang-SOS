# Break

> ## Por que? 

- A **estrutura** `for` é uma forma de gerenciarmos um conjunto de dados, e existem alguns métodos de ajudar esse gerenciamento, como o `break` e o `continue` . O **`continue` é quando precisamos encerrar o atual laço da  nossa estrutura de repetição** .


Exemplo :


```go
package main 

import "fmt"

func main () {
    for x := 0; x < 10; x++ {
        if x == 8 {
            continue
        }
        fmt.Println("Valor de x é", x)
    }
}

/*
    Valor de x é 0
    Valor de x é 1
    Valor de x é 2
    Valor de x é 3
    Valor de x é 4
    Valor de x é 5
    Valor de x é 6
    Valor de x é 7
    Valor de x é 9

*/
```

- Vimos que o **`break` irá encerra toda a estrutura do `for`** independentemente da condição inicialmente proposta ter sido atingida ou não, **o `continue` irá encerrar o atual laço do `for`**, com ele o `for` vai continar, normalmente, mas **para todo o laço em que ele estiver incluído a execução não será realizada .** 


