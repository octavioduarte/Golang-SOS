# Fake While 👀

## Fake ?

Já dissemos aqui que **Go possui poucas palavras reservadas**, uma das provas disso é a linguagem não tem **nativamente** o [While](https://en.wikipedia.org/wiki/While_loop), que é assim como o `For` uma estrutura de repetição presente em muitas linguagens . 

Mas existe uma solução pra essa "ausência" que se assemelha muito tanto na `sintaxe` quanto no `objetivo` usando o próprio `For` : 


```go
package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Println("Value of x", x)
		x++
	}
}

```

O equivalente em `Javascript` usando `While` seria: 

```javascript
var x = 0
while (x < 10) {
  console.log("Value of x", x)
  x++;
}
```