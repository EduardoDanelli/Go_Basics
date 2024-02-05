package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Número errado"
	}
}

func diaDaSemana2(numero int) string {

	var diaSemana string

	switch {
	case numero == 1:
		diaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaSemana = "Segunda"
	case numero == 3:
		diaSemana = "Terça"
	case numero == 4:
		diaSemana = "Quarta"
	case numero == 5:
		diaSemana = "Quinta"
	case numero == 6:
		diaSemana = "Sexta"
	case numero == 7:
		diaSemana = "Sabado"
	default:
		diaSemana = "Número errado"
	}
	return diaSemana
}

func main() {
	dia := diaDaSemana(7)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
