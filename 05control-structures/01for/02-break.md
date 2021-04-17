# Break

> ## Por que? 

- A **estrutura** `for` é uma forma de gerenciarmos um conjunto de dados, e existem alguns métodos de ajudar esse gerenciamento, como o `break` e o `continue` . O **`break` é quando precisamos encerrar nossa estrutura de repetição antes da condição imposta ser alcançada** .


Exemplo :


```go
package main 

import "fmt"

func main () {
    for x := 0; x < 10; x++ {
        if x == 8 {
            break
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
*/
```

- Repare que a condição imposta para a saída do `for` é que **x seja menor que 10**, mas há uma condicional (que é mais uma estrutura de controle que veremos em seguida) verificando **se x é igual a 8** e então usando o comando `break` **o compilador vai forçar o encerramento do fluxo do programa**, e mesmo 8 e 9 sendo menor que 10 não terão seus valores imprimidos no console .


