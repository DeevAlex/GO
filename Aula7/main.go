package main

import (
	"fmt"
)

func main() {

	// operadores

	// ARITMETICOS

	soma := 5 + 2
	subtracao := 5 - 2
	multiplicacao := 5 * 2
	divisao := 5 / 2

	// var numero1 int16 = 10
	// var numero2 int32 = 10

	// somaDiferentes := numero1 + numero2 -> nesse caso não vai funcionar, pois o GO ele não aceita operações com variaveis com tipos diferentes

	// ATRIBUIÇÃO

	var variavel1 string = "String 1"
	variavel2 := "String 2"

	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	// fmt.Println(somaDiferentes)

	// RELACIONAIAS

	fmt.Println(1 > 2)  // Maior que ...
	fmt.Println(1 >= 2) // Maior ou igual que ...
	fmt.Println(1 == 2) // Igual ...
	fmt.Println(1 <= 2) // Menor Igual ...
	fmt.Println(1 < 2)  // Menor que ...
	fmt.Println(1 != 2) // Diferente de ...

	// LOGICOS

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // As duas condicões tem que ser verdadeiras para retornar verdadeiro
	fmt.Println(verdadeiro || falso) // As uma condicão tem que ser verdadeira para retornar verdadeiro
	fmt.Println(!verdadeiro)         // inverte o valor, se verdadeiro ele retorna falso e vice-versa
	fmt.Println(!falso)              // inverte o valor, se falso ele retorna verdadeiro e vice-versa

	// UNARIOS

	numero := 10
	numero++     // numero = numero + 1
	numero += 15 // numero = numero + 15
	numero--     // numero = numero - 1
	numero -= 15 // numero = numero - 15
	numero *= 3  // numero = numero * 3
	numero /= 3  // numero = numero / 3
	numero %= 3  // numero = numero % 3
	fmt.Println(numero)

	// Não Existe operador ternario no GO, teria que fazer um if para implementa-lo

}
