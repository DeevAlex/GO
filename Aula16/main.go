package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2 // esses aqui serão o nome dado ali em cima para os return
	subtracao = n1 - n2
	// return soma, subtracao // antes
	return
}

func main() {

	// Funções com Retorno Nomeado

	soma, subtracao := calculosMatematicos(5, 5)

	fmt.Println(soma, subtracao)

}
