package main

import (
	"fmt"
)

// "Herança", Lembrando que no GO não existe a POO (Programação Orientada a Objetos), esse é um exemplo similar que chega mais proximo desse conceito

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float32
}

type estudante struct {
	pessoa    // pegando os atributos acima, não repita o tipo
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"José", "Lima D.", 29, 1.75}

	e1 := estudante{p1, "Ciencia da Computação", "UnG"}

	fmt.Println(e1)
	fmt.Println("Nome: " + e1.nome) // pegando um valor acessando o atributo

}
