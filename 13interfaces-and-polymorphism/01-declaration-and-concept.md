> ## Interfaces e Polimorfismo 🔥

Ao estudarmos sobre [Métodos](https://github.com/octavioduarte/Golang-SOS/blob/master/12methods/01-concept.md) na sessão anterior, podemos concluir que `Go` tem diferentes formas de chegar em uma solução se o compararmos com linguagens orientadas a objetos . 

Isso fica claro quando estudamos também sobre `Interfaces` e `Polimorfismo` que são conceitos que devem ser estudados de maneira conjunta pois são bem relacionados um ao outro em `Go` .


> Interfaces 

Se estudarmos sobre `Interfaces` em  `Typescript`, iremos concluir que elas são usadas para criarmos a assinaturas de métodos, e toda classe que implementa este método deve respeitar sua assinatura . 

Exemplo : 

```typescript 
interface HumanActions {
    greet: (who: string) => string
    sleep: (isNight: boolean) => string
    wakeUp: () => string
    work: (place: string) => string
}

class Human implements HumanActions {

    greet(who: string): string {
        return `Hello ${who}`
    }

    sleep(isNight: boolean): string {
        return `I'm ${!isNight && 'not '} sleeping`
    }

    wakeUp(): string {
        return "I woke up"
    }

    work(place: string): string {
        return `I'm work in ${place}`
    }
}
```

Na prática podemos concluir que `interfaces` nada mais é do que **um conjunto de assinaturas de métodos para padronizarmos as classes que vão implementa-la** .

Em `Go` este conceito não é tão distante do cenário `Typescript`, na prática `interfaces` em `Go` também são um conjunto de assinaturas de métodos na qual todo método que a implementa deve respeitar **explicita e rigidamente** sua assinatura . 


> Polimorfismo 

Certo, agora já entendemos que **`Interfaces` são um conjunto de assinaturas de métodos na qual todo método que a implementa deva respeitar de forma **explicita e rigida** sua assinatura** . Agora é interessante entender que um dois métodos podem ter assinaturas identicas e ao mesmo tempo propósitos completamente diferentes, ou seja, diferentes métodos que respeitam a assinatura de uma interface, mas tem contextos, propósitos e consequentemente um código diferente . 

Podemos ver isso neste exemplo : 


```go
package main

import "fmt"

// Interface log diz que o método que quem implementa-la, deve ter um método sem parâmetro e sem retorno, chamado say
type logs interface { 
	say()
}

// Struct human 👩
type human struct {
	name   string
	pharse string
}

// Método say() vinculado ao struct human
func (h human) say() {
	fmt.Println("The human", h.name, "is speakin", h.pharse)
}

// Struct dog 🐕
type dog struct {
	color  string
	pharse string
}

// Método say() vinculado ao struct dog --> Com mensagem diferentes e propósitos diferentes
func (d dog) say() {
	fmt.Println("The dog is speakin", d.pharse)
}


func handleLogs(l logs) {
	l.say()
}

func main() {
    // instância do struct human
	vinnyAppice := human{name: "Viny Appice", pharse: "Thanks Ronnie !!!"}

    // instância do struct dog
	bob := dog{color: "black", pharse: "AU AU AU"}


    // Execução do método, isso requer bastante atenção, neste momento estamos executando métodos de diferentes instâncias,
    // com diferentes propósitos e contextos com um único trecho de código. Isso só é possível pois ambos os structs respeitam
    // de forma rigida e explicita a interface `logs`
	handleLogs(vinnyAppice)
	handleLogs(bob)


    // Em programação polimorfismo é a prática de invocar métodos que possuem a mesma assinatura, mas comportam-se de maneiras
    // diferentes
}

// The human Viny Appice is speakin Thanks Ronnie !!!
// The dog is speakin AU AU AU
```