> ## Comma ok


**Ao acessarmos uma chave inexistente de um `Map`, nosso retorno não é um erro, mas sim um 0**, isso pode causar um problema, pois seria impossível diferenciar se aquele 0 corresponde de fato a um valor do `Map`, ou se aquela chave não existe . 

É para isso que serve o `comma ok`, ele basicamente extrai do `Map` um valor `booleano`, na qual `true` significa que aquela chave corresponde de fato a um valor e `false` que aquela chave não existe no `Map` : 



```go
package main

import "fmt"

func main() {
	members := map[string]int{
		"Robert Johnson": 1938,
		"Janis Joplin":   1970,
		"Jimi Hendrix":   1970,
		"Jim Morrison":   1971,
		"Kurt Cobain":    1994,
		"Fake Member":    0,
	}

	valueOzzy, okOzzy := members["Ozzy Osbourne"]
	fmt.Println(valueOzzy, okOzzy)

	valueFakeMember, okFakeMember := members["Fake Member"]
	fmt.Println(valueFakeMember, okFakeMember)
}
// 0 false
// 0 true
```