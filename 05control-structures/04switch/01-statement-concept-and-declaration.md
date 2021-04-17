# Switch Statement

## Declaração 

 Assim como o `if` o **`switch` também é uma estrutura de controle condicional**, ou seja, vai analisar uma condição e tomar uma determinada ação a partir do resultado dessa análise. 



```go
package main

import "fmt"

func main() {
	country := "Brasil"

	switch country {
	case "China":
		fmt.Println(country, "fica na Ásia")
	case "Brasil":
		fmt.Println(country, "fica na America do Sul")
	case "Russia":
		fmt.Println(country, "fica na Europa e Ásia")
	}
}

// Brasil fina na America do Sul

```

## Por que o Switch e não o If (e vice-versa)

É uma dúvida comum, pois ambas tem um propósito bem similar o que pode gerar uma confusão na hora de escolher qual o melhor ,

**Opinião pessoal**

Gosto mais da sintaxe do `switch` e acho ele mais objetivo, mas assim como ele analisa condições para tomar uma ação eu também faço o mesmo . 

- Uso o Switch se : 

* Todos os possíveis resultados pra minha condicional são previsiveis, o `switch` até possibilita um caminho para um resultado não previsivel com o `default`, mas acho que perde um pouco do sentido usar o `switch` senão tenho mapeado as possibilidades .

* Os possíveis resultados se resumem a valores primitivo . 

**Condicionais que não estão nesse escopo devem ser analisadas com `if`**