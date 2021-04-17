# Variáveis

> ## Declaração 

```go
var name string = "Janis Joplin"
var good_singer bool = true
var age int = 27
```

> ## Notas

* Apesar de `Go` ser uma **linguagem fortemente tipada**, ele permite a declaração
de variáveis sem tipagem explicita . O compilador se encarrega da atribuíção de tipagem baseado
no valor atribuído . 

```go
var city = "Port Arthur"
```

- É possível declarar variáveis sem atribuição, o Go se encarregará de atribuir o equivalente do
valor nulo 


```go
var document int;   // Será atribuído   ==> 0
var country string; // Será atribuído   ==> ""
var logged  bool;   // Será atribuído   ==> false
```