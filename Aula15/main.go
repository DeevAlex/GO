package main

import (
	"fmt"
	"time"
)

func main() {

	// Loops

	// i := 0

	// for
	// for i < 10 {
	// 	i++;
	// 	fmt.Println("i:", i)
	// 	time.Sleep(time.Microsecond)
	// }

	// for init
	// for j := 0; j < 10; j+=2 {
	// 	fmt.Println("j:", j)
	// 	time.Sleep(time.Microsecond)
	// }

	// for com a clausula range - mais usado quando vamos iterar um objeto

	nomes := [3]string{"Alex", "Marina", "Luiza"}

	// o primeiro parametro sempre vai ser o indice e o segundo o valor, para querer um parametro apenas coloque _
	for indice, valor := range nomes {
		fmt.Println("indice e nome:", indice, valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println((indice + 1), string(letra)) // na letra que ele pega o codigo da tabela ascii, para pegar a letra é so converter para string
	}

	usuario := map[string]string{
		"nome":      "Alex",
		"sobrenome": "Caldeira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// loop infinito

	for {
		fmt.Println("INFINITO!!!!")
		time.Sleep(time.Second)
	}

	// não tem como iterar sobre struct
	// Ex:

	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"Zé", "Junior"}

	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }

}
