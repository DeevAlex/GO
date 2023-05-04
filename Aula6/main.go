package main

import (
	"fmt"
)

// Funções

//   nome  parametros/tipo   tipo retorno
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// nesse caso os parametros vao ser do mesmo tipo -> int
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao // o go permite mais de um retorno, definimos os tipos de retorno ali em cima no (int8, int8)
}

func main() {

	// atribuindo uma função a uma variavel

	var f = func() {
		fmt.Println("Minha função na variavel f")
	}

	var retornaTexto = func(txt string)  {
		fmt.Println(txt)
	}

	soma := somar(5, 5)

	fmt.Println(soma)
	f()
	retornaTexto("Olá Mundo")

	// retorno soma / retorno subtracao
	resultadoSoma, resultadoSubtracao := calculosMatematicos(5, 7)
	// se não quisermos um retorno colocamos _ 
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSubtracao)
}


