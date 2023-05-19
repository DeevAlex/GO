package main

import (
	"fmt"
)

func main() {

	// Estrutura de Controle

	// 1ª forma
	if (5 == 5) {
		fmt.Println("5 é igual a 5")
	} else {
		fmt.Println("5 não é igual a 5")
	}

	// 2ª forma

	numero := 50

	if numero > 25 {
		fmt.Println("O número é maior que 25")
	} else {
		fmt.Println("O número é menor que 25")
	}

	// if init

	// criando e checkando uma condição
	if outroNumero := numero; outroNumero > 25 {
		fmt.Println("O número é maior que 25")
	} else {
		fmt.Println("O número é menor que 25")
	}

}
