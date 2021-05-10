> ## Defer

O `defer` é uma das **poucas palvras reservadas da linguagem `Go`**, seu objetivo é **determinar que a linha de código na qual ele se encontra será a ultima coisa a ser executada antes do retorno** . 


```go
package main

import "fmt"

func main() {
	defer fmt.Println(execCount()) // Por conta do defer, esta linha será executada por ultimo.
	fmt.Println("Janis Joplin passed away with: ") 
}

func execCount() int {
	return 1970 - 1943
}

// Janis Joplin passed away with: 
// 27

```