package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")
	var user usuario
	user.idade = 25
	user.nome = "Gaba"
	fmt.Println(user)
	toString(user)

	end := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Jonas", 24, end}
	fmt.Println(u2)
	toString(u2)

	u3 := usuario{nome: "Bigode"}
	fmt.Println(u3)
	toString(u3)
}

func toString(user usuario) {
	fmt.Println(
		" Nome:", user.nome, "\n",
		"Idade:", user.idade, "\n",
		"Endereco:", user.endereco.logradouro, user.endereco.numero)
}
