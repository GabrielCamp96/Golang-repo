package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // Retira 1 do waitGroup (2 - 1)
	}()

	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() // Retira 1 do waitGroup (1 - 1)
	}()

	go func() {
		escrever("GO Routine 3!")
		waitGroup.Done() // Retira 1 do waitGroup (1 - 1)
	}()

	go func() {
		escrever("GO Routine 4!")
		waitGroup.Done() // Retira 1 do waitGroup (1 - 1)
	}()

	waitGroup.Wait()
}
