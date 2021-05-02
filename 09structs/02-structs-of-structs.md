> ## Structs of Structs


É possível combinarmos diferentes `Structs` dentro de um único com a seguinte sintaxe



```go
package main

import "fmt"

type singerInfo struct {
	born int
	dead bool
	name string
}

type band struct {
	active bool
	singerInfo
	name    string
	country string
	year    int
}

func main() {
	dio := band{
		active: false,
        // Isso mesmo, devemos repetir 😉
		singerInfo: singerInfo{
			name: "Ronnie James Dio",
			born: 1942,
			dead: true,
		},
		name:    "DIO",
		country: "EUA",
		year:    1982,
	}

	fmt.Println(dio)
}



// {false {1942 true Ronnie James Dio} DIO EUA 1982}
```