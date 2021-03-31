# Iota

> ## Conceito 


- Iota é o termo técnico que Go usa para criar uma sequência de números inteiros sucessivos


```Code 
package main

import "fmt"

const (
	C0 = iota
	C1 = iota
	C2 = iota
)

func main() {
	fmt.Println(C0, C1, C2)
}


# Output

// 0 1 2
```

