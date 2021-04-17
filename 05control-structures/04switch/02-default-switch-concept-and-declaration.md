# Switch Default

## Sobre 

É muito similar ao `switch statement`, **o que vai diferir é que com ele o parâmetro de análise passado ao switch pode ser omitido e o compilador vai se encarregar de atribuir um `true` como parâmetro de análise e o primeiro `case` que satisfazer a análise com esse resultado será executado** 



```go
package main

import "fmt"

func main() {

	userProfile := "admin"

	// O parâmetro de análise não foi passado
	switch {
	case userProfile == "root":
		fmt.Println("Usuário é root")
	case userProfile == "admin":
		fmt.Println("Usuário é admin")
	case userProfile == "guest":
		fmt.Println("Usuário é guest")
	}
}


// Usuário é admin

```