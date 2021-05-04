> ## Functions with params


Como ja observamos `Go` não tem muitas particularidades na sintaxe de suas funções se formos comparar com outras **linguagens tipadas .** A passagem de **parâmetros**  e a forma como recebemos isso como **argumentos** em nossas funções é um bom exemplo disso .   

Nas **funções** escrevemos todos os argumentos a serem recebidos com suas respectivas tipagens : 

```go

func functionWithParams (param1 int, param2 string, param3 bool) {
    // code ...
}

```

> Se a tipagem for a mesma para todos podemos escreve-la apenas uma vez

```go
func funcWithStringParams (name, lastname string) {
    // code ...
}
```