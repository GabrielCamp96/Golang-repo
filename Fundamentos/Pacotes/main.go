package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Escrever registra uma mensagem na tela
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("teste@teste.com.br")
	fmt.Println(err)
	// auxiliar.escrever2() Não pode se referir a ela, pois é um método privado (inicia com letra minúscula)
}
