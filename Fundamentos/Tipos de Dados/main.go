package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100 // int8, int16, int32, int64
	fmt.Println(numero)

	numeroInferido := 100000000000000
	fmt.Println(numeroInferido)

	// uint - unsigned int

	// alias
	var numero3 rune = 123456 // int32 = rune
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12333333.45
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 's'
	fmt.Println(char)

	var texto string
	fmt.Println(texto) // Valor 0 - Valor inicial

	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := false
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro)

	var erroTratado error = errors.New("erro tratado")
	fmt.Println(erroTratado)
}
