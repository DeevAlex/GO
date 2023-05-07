package main

import (
	"fmt"
)

func main() {

	// Ponteiros

	var variavel1 int = 10
	// para haver alteração tem que ser antes
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// nesses casos aqui ele só vai alterar o valor de uma variavel, apos colocar uma variavel para receber o valor de outra não havera alteração
	variavel1++

	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA, no caso do ponteiro pode alterar o valor da variavel que altera a outra varivael também, poruq ela só muda o valor não o endereço de memoria

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // endereço de memoria

	fmt.Println(variavel3, *ponteiro) // esse *<variavel>, mostra o valor que esta nessa memoria, chamada de desreferenciação

}
