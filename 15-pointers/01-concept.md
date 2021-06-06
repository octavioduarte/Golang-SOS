> ## Funções Anônimas ➡️  ⏲️

`Ponteiro` é um valor que armazena o endereço de memória de uma variável, é importante lembrar que **o valor das variáveis são armazenados em disco** e o ponteiro é o responsável por armazenar o endereço do disco na qual o valor daquela variável esta alocado . 


Em `Go` para que tenhamos acesso ao endereço da memória na quel o valor de nossas variáveis estão alocados, usamos a seguinte sintaxe 


```go

import "fmt"

func main() {

	x := 10
	addressMemoryOfX := &x // Basta colocarmos um "E comercial" antes do nome da variável

	fmt.Println("The address memory of variable x is ", addressMemoryOfX)
}

// The address memory of variable x is  0xc000018090
```