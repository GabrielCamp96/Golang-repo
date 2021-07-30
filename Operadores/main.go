package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	mult := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, mult, resto)

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println("-----------------")
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 20
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 10
	fmt.Println(numero)
	numero %= 2
	fmt.Println(numero)
}
