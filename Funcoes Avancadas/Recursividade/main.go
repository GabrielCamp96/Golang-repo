package main

import "fmt"

func fib(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fib(posicao-2) + fib(posicao-1)
}

func main() {
	seq := uint(12)

	for i := uint(1); i <= seq; i++ {
		fmt.Println(fib(i))
	}
}
