package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada, resultado será calculado")
	fmt.Println("Entrando na funcao para ver se o aluno foi aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Defer")
	defer funcao1() // DEFER = adiar
	funcao2()
	fmt.Println(alunoEstaAprovado(6, 6.5))
}
