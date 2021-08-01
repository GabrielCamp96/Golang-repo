package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // variavel existe apenas no escopo de execução do IF INIT
		fmt.Println("É maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("É menor a -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
