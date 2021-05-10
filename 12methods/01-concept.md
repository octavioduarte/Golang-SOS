> ## Methods 🔥

`Methods` é um conceito em `Go` mais simples de ser compreendido quando conhecemos sobre [POO](https://pt.wikipedia.org/wiki/Programa%C3%A7%C3%A3o_orientada_a_objetos)(Programção Orientada a Objetos) . 

Em POO, temos o conceito de `Classes` e `atributos` e `métodos` vinculados a ela, na prática, isso simplesmente se resume em criar uma contexto com todos os métodos e dados pertencentes a esse contexto em um unico lugar .


Um exemplo em `Javascript` : 

```javascript
class Dog {
  constructor(name, age, color) {
    this.name = name
    this.age = age
    this.color = color
  }

  printAge() {
    console.log(`${this.name} has ${this.age} years`)
  }

  printColor() {
    console.log(`The color of ${this.name} is ${this.color}`)
  }

  printBark() {
    console.log(`${this.name} is talking AU AU AU`)
  }
}
const dog1 = new Dog('Guga', 18, 'White')
dog1.printAge()

const dog2 = new Dog('Jadi', 20, 'Black')
dog2.printBark()

const dog3 = new Dog('Bob', 4, 'Gray')
dog3.printColor()

// Guga has 18 years
// Jadi is talking AU AU AU
// The color of Bob is Gray
```

No exemplo acima temos a `classe` Dog e nela temos os `métodos` **printAge**, **printColor** e **printBark**, ambos os métodos são executados imprimindo os parâmetros passados no `construtor` da `classe`, que são **name**, **age** e **color** .

( Repare que apesar de terem propósitos diferentes (seja `métodos` ou `atributos`) estamos falando do mesmo contexto, no caso cachorros.) 

`Go` tem uma forma diferente de lidar com esse paradigma, usando `métodos` . 

Na prática criamos `structs` e vinculamos a esses structs métodos e uma vez vinculados, toda instância deste struct **(e somente a instância)** tera acesso a este método . 

O código em `Go` equivalente ao código em `Javascript` acima, seria algo desse tipo : 


```go
package main

import "fmt"

type dog struct {
	name  string
	age   int
	color string
}

func (d dog) printAge() {
	fmt.Println(d.name, "has", d.age, "years")
}

func (d dog) printColor() {
	fmt.Println("The color of", d.name, "is", d.color)
}

func (d dog) printBark() {
	fmt.Println(d.name, "is talking AU AU AU")
}

func main() {
	dog1 := dog{"Guga", 18, "White"}
	dog2 := dog{"Jadi", 20, "Black"}
	dog3 := dog{"Bob", 4, "Gray"}

	dog1.printAge()
	dog2.printBark()
	dog3.printColor()
}

// Guga has 18 years
// Jadi is talking AU AU AU
// The color of Bob is Gray

```