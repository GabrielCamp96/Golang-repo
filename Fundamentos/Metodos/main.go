package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) getNome() string {
	return u.nome
}

func (u usuario) getIdade() uint8 {
	return u.idade
}

func (u *usuario) setNome(nome string) {
	u.nome = nome
}

func (u *usuario) setIdade(idade uint8) {
	u.idade = idade
}

func main() {
	fmt.Println("Métodos")

	u := usuario{"Joao", 35}
	u2 := usuario{"Carlos", 40}

	fmt.Println(u.getNome(), u.getIdade())
	fmt.Println(u2.getNome(), u2.getIdade())

	u.setIdade(25)
	fmt.Println(u.getNome(), u.getIdade())

	u2.setNome("Davi")
	fmt.Println(u2.getNome(), u2.getIdade())
}
