package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	// CONCORRENCIA != PARALELISMO

	go escrever("Olá mundo!") //go routine
	escrever("Programando em GO!")
}
