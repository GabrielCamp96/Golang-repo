package main

import "fmt"

func soma(numeros ...int) (soma int) {
	for _, n := range numeros {
		soma += n
	}
	return
}

func main() {
	fmt.Println("Funções variáticas")

	n := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println("Soma =", n)
}
