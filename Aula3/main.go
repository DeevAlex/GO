package main

import (
	"fmt"
	"modulo/auxiliar"
)


func main() {
	fmt.Println("Escrevendo do Arquivo Main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("alexdematoscaceresmail")

	fmt.Println(erro)
}

// comando para pegar pacotes externos: go get <endereço do pacote>
// endereço de exemplo: github.com/badoux/checkmail

// para usar chamar uma função de um pacote tem que ser o nome que está na ultima /<pacote que sera usado para chamar a função>
// nulo do go: <nil>

// para remover todas as dependencias que não estão em uso: go mod tidy