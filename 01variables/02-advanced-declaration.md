# Variáveis

> ## Declaração 

- É possível criar diversas variáveis simultaneamente desde que todas tenham um único tipo : 

```code 
    var name, job, city string

    /* 
        Neste caso teríamos 3 diferentes variáveis com o valor 
        default de strings (" ") atribuídos a elas
    */
```


- Outra forma que também deixa o código mais semântico : 

```code 
    var (
        item string
        price float64
        quantity int
    )
```


- Declaração com atribuição sem especificar contexto e tipagem : 


```code 
	message := "Hello"
```

- O exemplo acima só pode ser usado em um codeblock . Lembrando que em Go há dois principais escopos : 

* package-level-scope
* codeblock

```code
package {

    // package-level-scope

    func main() {
        // codeblock
    }
}
```

- E quando desejar utilizar declaração + atribuição de uma variável, só é possível em um codeblock . 

