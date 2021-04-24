> ## Append dos valores de um slice em outro slice

Não há uma grande diferença na sintaxe ao anexar uma `slice` a uma outra, em resumo temos que resolver apenas um detalhe relacionado a tipagem .  



```go
package main

import "fmt"

func main() {
	start := []string{"Tô", "ouvindo", "alguém", "gritar", "meu", "nome"}
	complement := []string{"parece", "mano", "meu", "é", "voz", "de", "homem"}

	start = append(start, complement...)

	fmt.Println(start)
}
```

**Caso seja omitido os 3 pontos, receberemos um erro** informando que estamos tentando incluir em na variavel `start`, cujo o tipo é `[]string`, ou seja, um slice com `strings` agrupadas, um novo `slice` que também possui outras `strings` agrupadas . 

<br />

É **importante** entender que : 

```go
[]string !== string

[]string ==> {"Uma", "lista", "de", "strings"}
string   ==> "Uma unica string"
```

Ou seja um **`array` de `strings`**  só pode receber `append` com `strings` . 

É o que os 3 pontos nos ajudam a fazer, ele basicamente diz ao compilador para considerar apenas os valores de um `slice` e desde que os valores respeitem a tipagem do `slice` a ser alterado, o código funcionará normalmente . 