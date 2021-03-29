# Tipos

> ## Tipos Customizados 

- Em Go é possível criar uma tipagem customizada :


```code 
    package main

    import "fmt"

    type myCustomString string
    type myCustomInt int

    var (
        var_str myCustomString = "Woodstock"
        var_int myCustomInt    = 1969
    )

    func main() {
        fmt.Printf("Type of %v is %T\n", var_str, var_str)
        fmt.Printf("Type of %v is %T\n", var_int, var_int)
    }

    #  Output
    
    // Type of Woodstock is main.myCustomString
    // Type of 1969 is main.myCustomInt
```


