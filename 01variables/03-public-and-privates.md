# Variáveis

> ## Publicas e Privadas 

- **Go é uma linguagem com poucas palavras reservadas**, em boa parte dos casos basta uma simples 
customização no código e já é possível obter o resultado esperado . 

- É o caso de variáveis `públicas` e `privadas`. 

**Em Go a visibilidade de uma variável está condicionada ao  primeiro caractere do nome da variável** . 

* Se for maiusculo estamos informando que aquela variável é pública e pode ser usada fora do pacote
* Se for minusculo estamos informando que aquela variável é privada e não pode ser usada fora do pacote

```go 
var (
	PublicVar, privateVar string
)
```