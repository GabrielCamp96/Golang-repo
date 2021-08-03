package main

import "fmt"

func fib(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fib(posicao-2) + fib(posicao-1)
}

func worker(tarefas <-chan uint, resultados chan<- uint) {
	for tarefa := range tarefas {
		resultados <- fib(tarefa)
	}
}

func main() {
	tarefas := make(chan uint, 45)
	resultados := make(chan uint, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- uint(i)
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	close(resultados)
}
