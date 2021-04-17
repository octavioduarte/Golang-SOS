# Variáveis

> ## Declaração 

* É possível criar diversas variáveis simultaneamente desde que todas tenham um único tipo : 

```go 
    var name, job, city string

    // Neste caso teríamos 3 diferentes variáveis com o valor default de strings (" ") atribuídos a elas
```


* Outra forma que também deixa o código mais semântico : 

```go 
    var (
        item string
        price float64
        quantity int
    )
```


- Declaração com atribuição sem especificar contexto e tipagem : 


```go 
	message := "Hello"
```

- O exemplo acima só pode ser usado em um codeblock . Lembrando que em Go há dois principais escopos : 

* package-level-scope
* codeblock

```go
package {

    // package-level-scope

    func main() {
        // codeblock
    }
}
```

**E quando desejar utilizar declaração + atribuição de uma variável, só é possível em um codeblock .**

