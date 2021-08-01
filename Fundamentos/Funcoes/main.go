package main

import "fmt"

func main() {
	fmt.Println(somar(2, 3))

	var f = func(txt string) string {
		fmt.Println("Função f:", txt)
		return txt
	}

	f("Parametro")

	resultado := f("Texto da funcao 1")
	fmt.Println(resultado)

	retorno1, retorno2 := calculadora(10, 15)
	fmt.Println(retorno1, retorno2)

	retorno3, _ := calculadora(10, 15)
	fmt.Println(retorno3)

	_, retorno4 := calculadora(10, 15)
	fmt.Println(retorno4)
}

func somar(x int8, y int8) int8 {
	return x + y
}

func calculadora(x, y int8) (int8, int8) {
	return x + y, x - y
}
