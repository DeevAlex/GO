package main

import (
	"fmt"
	"modulo/auxiliar"
)


func main() {
	fmt.Println("Escrevendo do Arquivo Main")
	auxiliar.Escrever()
}

// 1º cd <pasta>
// 2º go mod init <nome do modulo>
// 3° go build - para compilar o código
// 4º <caminho ou pasta> que ira executar esse código compilado
// se quiser criar mais de um pacote tem que criar mais de uma pasta
// quando começar uma função com letra maiuscula indica que ela será publica e privada caso seja minuscula