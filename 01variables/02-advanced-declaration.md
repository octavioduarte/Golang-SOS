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


- Criação com atribuição sem especificar contexto e tipagem : 


```code 
	message := "Hello"
```
