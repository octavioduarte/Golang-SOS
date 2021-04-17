# Iota

> ## Operações 


- Agora sabemos que o Iota nos permite criar com uma sequência numérica de forma mais simples, mas é importante entender também que ele possibilita criar operações que vão permitir que possamos customizar como essa sequência será iniciada .


- Como ?


- Basta que passemos uma operação cada item da sequência do Iota


```Code 
package main

import "fmt"

const (
	C1 = (iota + 1) * 5     //  ( 0 <--- iota  + 1 ) *  5 = 5
	C2 = (iota + 2) * 10    //  ( 1 <--- iota  + 2 ) * 10 = 30
	C3 = (iota + 3) * 20    //  ( 2 <--- iota  + 3 ) * 20 = 100
)

func main() {
	fmt.Println(C1, C2, C3)
}

# Output

// 5 30 100
```

