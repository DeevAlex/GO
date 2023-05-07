package main

import (
	"fmt"
)

// sintaxe do structs
type usuario struct {
	nome string
	idade int
	endereco endereco // struct dentro de struct
}

type endereco struct {
	logradouro string
	numero int
}

func main() {

	// structs -> coleção de campos

	var u1 usuario // valor padrão serão o valor 0 de cada tipo

	u1.nome = "Alex"
	u1.idade = 19

	enderecoExemplo := endereco{"Rua Nizal", 77}

	u2 := usuario{"Jonas", 21, enderecoExemplo}

	// para preencher um campo é so colocar a chave e o valor
	u3 := usuario{nome: "Lucas"}

	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)

}
