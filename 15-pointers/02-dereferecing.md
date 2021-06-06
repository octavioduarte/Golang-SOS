> ## Referência a ponteiros

Existe uma forma de acessarmos o valor de um endereço de memória 

```go
package main

import "fmt"

func main() {

	x := 10
	addressMemoryOfX := &x
	contentOfAddressOfMemoryOfX := *addressMemoryOfX

	fmt.Println("The content of", addressMemoryOfX, "is", contentOfAddressOfMemoryOfX)
}
// The content of 0xc00018c000 is 10
```