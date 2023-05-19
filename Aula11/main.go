package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Array e Slices

	var array1 [5]int // todos os dados deveram ser do tipo da lista, caso o array esteja vazio o GO preencherá com o valor padrão do tipo da lista

	// atribuindo valores pelo indice
	array1[0] = 5
	array1[1] = 6
	array1[2] = 7
	array1[3] = 8
	array1[4] = 9

	// usando a referencia de tipos

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	array2[4] = "Olá Mundo"

	// Array mais flexivel

	array3 := [...]int{1, 2, 3, 4, 5} // esses tres pontos vai ser a quantidade de items que tem dentro do array sem fixar um valor

	// array3[8] = 5, isso não é permitido, pois os [...] não deixa o array com um tamanho dinamico

	// Slices

	slice := []int{2, 4, 5, 6, 8, 10, 12, 14, 16, 18, 20} // slice não é um array ele se parece com um, o slice seria como uma fatia de um array

	slice = append(slice, 22) // essa função faz com que adicionamos um numero a um slice, sem passar o indice

	// criando uma fatia entre intervalos

	slice2 := array2[1:4]

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(reflect.TypeOf(slice)) // retorna o tipo de alguma coisa que está dentro dos ()
	fmt.Println(reflect.TypeOf(array1))

	array2[1] = "Posição Alterada" // os slices funcionam como o ponteiro, se alterar um ele altera o outro
	fmt.Println(slice2)

	// Arrays Internos

	//              tipo  tamanho capacidade maxima
	slice3 := make([]float32, 10, 11) // esse metodo make que vai alocar espaço na memoria para determinada coisa, como um slice por exemplo
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // metodo len é o para ver o tamanho
	fmt.Println(cap(slice3)) // metodo cap é o para ver a capacidade

	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	slice3 = append(slice3, 1)
	fmt.Println(slice3)

	// o Go quando vê que o slice vai passar a capacidade ele cria um outro array e dobra o tamanho dele

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// aqui caso não passe a capacidade maxima ele pega o tamanho que foi passado e usa como capacidade desse slice
	slice4 := make([]float32, 5)

	slice4 = append(slice4, 10)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// resumindo slice é uma lista sem tamanho fixo e array é com tamanho fixo

}
