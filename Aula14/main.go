package main

import (
	"fmt"
)

// assim ou
func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado-feira"
	default:
		return "Forneça um dia valido de 1 a 7"
	}
}

// assim
func diaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // é clausula, ele executa o codigo e parte para a proxima condição
		// e não precisa usar a clausula break
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado-feira"
	default:
		diaDaSemana = "Forneça um dia valido de 1 a 7"
	}

	return diaDaSemana

}

func main() {

	// Switch

	fmt.Println(diaSemana(3))

	fmt.Println("")

	fmt.Println(diaSemana2(1))

}
