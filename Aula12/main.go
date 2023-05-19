package main

import (
	"fmt"
)

func main() {

	// maps

	//   [tipo das chaves]tipos dos valores
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Roberto",
		"idade":     "25",
	}

	// map aninhando, muitos maps ficam desorganizado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Roberto",
			"segundo":  "Campos",
		},
		"faculdade": {
			"campus":      "Guarulhos",
			"instituição": "Cruzeiro do Sul",
			"curso":       "Ciencia da Computação",
		},
		"linguagens": {
			"Linguagem 1": "Golang",
			"Linguagem 2": "Java",
			"Linguagem 3": "JavaScript",
		},
	}

	fmt.Println(usuario["nome"]) // acessando o chave pelo [] e dentro do [] coloca o que você quer pegar
	fmt.Println(usuario2["faculdade"]["campus"])
	fmt.Println(usuario2)
	delete(usuario2, "linguagens") // deleta a chave ... do map ...
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{ // criando uma chave para o map ...
		"nome": "Sagittarius",
	}
	fmt.Println(usuario2)

}
