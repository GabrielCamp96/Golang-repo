package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplex(escrever("Olá mundo!"), escrever("Programando em GO!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplex(entrada1, entrada2 <-chan string) <-chan string {
	saida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-entrada1:
				saida <- mensagem
			case mensagem := <-entrada2:
				saida <- mensagem
			}
		}
	}()

	return saida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
