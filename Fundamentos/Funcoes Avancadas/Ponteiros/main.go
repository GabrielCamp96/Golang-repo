package main

import "fmt"

func inverterSinal(numero *int) {
	*numero *= -1
}

func main() {
	fmt.Println("Função com ponteiro")

	numero := 20
	inverterSinal(&numero)
	fmt.Println(numero)
}
