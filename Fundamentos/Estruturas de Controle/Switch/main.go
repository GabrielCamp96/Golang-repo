package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Numero fora do range 1-7"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		// fallthrough // Força a próxima condição
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Numero fora do range 1-7"
	}

	return diaDaSemana
}
func main() {
	fmt.Println("Switch")

	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(8))
	fmt.Println(diaDaSemana2(1))
}
