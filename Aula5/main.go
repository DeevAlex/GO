package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de dados

	// int, int8, int16, int32, int64 - tipo inteiro, o que muda é a capacidade de bits
	// int - > usa a arquitetura como base Ex: windows 64x bit ele usará o int64 por padrão caso não seja explicito na declaração

	var numero int16 = 100 // explicito
	numero1 := -1000       // implicita

	// uint -> unsygned int, ele não aceita numeros negativos

	var numero2 uint = 1

	// alias

	// INT32 = RUNE
	var numero3 rune = 15

	// INT8 = BYTE
	var numero4 byte = 15

	// float

	// float32
	var numeroReal1 float32 = 1.5

	// float32
	var numeroReal2 float32 = 1000.5

	numeroReal3 := 1.4

	// Strings

	var str string = "Texto"
	str2 := ""

	char := '6' // retorna o numero decimal em que esse numero/letra está na tabela ASCII

	// boolean

	var booleano bool = true

	// error

	var erro error = errors.New("Erro Interno") // atribuir um erro a essa variavel

	// Valor 0 das variaveis

	// var texto string - valor 0: string vazia
	// var valor0int int - valor 0: 0
	// var valor0float float32 - valor 0: 0
	// var valor0rune rune - valor 0: 0
	// var valor0byte byte - valor 0: 0
	// var valor0bool bool - valor 0: false
	// var valor0error error - valor 0: <nil>

	fmt.Println(numero)
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numeroReal1)
	fmt.Println(numeroReal2)
	fmt.Println(numeroReal3)
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(char)
	fmt.Println(booleano)
	fmt.Println(erro)

}

// Tabela ASCII -> https://www.matematica.pt/util/resumos/tabela-ascii.php
