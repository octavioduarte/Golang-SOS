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


