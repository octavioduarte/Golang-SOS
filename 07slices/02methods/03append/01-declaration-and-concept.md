> ## Append

Usamos o [`append`](https://gobyexample.com/slices) para criar incluir novos dados em um `Slice`.




```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4}
    slice = append(slice, 5, 6, 7)

    fmt.Println(slice)
}


// [1 2 3 4 5 6 7]
```
Sua sintaxe se resume a 

```go
slice_a_ser_alterado = append(slice_a_ser_alterado, novos valores)
```