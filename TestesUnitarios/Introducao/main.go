package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	fmt.Println(enderecos.TipoDeEndereco("Rua dos bobos"))
	fmt.Println(enderecos.TipoDeEndereco("Viela dos bobos"))
}
