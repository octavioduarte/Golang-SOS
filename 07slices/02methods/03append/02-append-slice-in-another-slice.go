package main

import "fmt"

func main() {
	start := []string{"Tô", "ouvindo", "alguém", "gritar", "meu", "nome"}
	complement := []string{"parece", "mano", "meu", "é", "voz", "de", "homem"}

	start = append(start, complement...)

	fmt.Println(start)
}
