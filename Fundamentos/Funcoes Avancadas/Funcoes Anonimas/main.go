package main

import "fmt"

func main() {
	fmt.Println("Funções anonimas")

	s := func(text string) string {
		fmt.Println("Hello world!", text)
		return fmt.Sprintf("Recebido -> %s", text)
	}("Passando parametro")

	fmt.Println(s)

}
