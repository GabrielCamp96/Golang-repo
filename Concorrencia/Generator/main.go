package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Encapsula a chamada de uma Go routine e devolve um canal para utilizar as informações
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
