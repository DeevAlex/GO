package main

import (
	"fmt";
)


func main() {

	// variavel

	var variavel1 string = "Variavel 1"; // explicito
	variavel2 := "Variavel 2" // implicito

	var (
		variavel3 string = "Variavel 3"
		variavel4 string
	)

	// ou

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	// constante

	const constante1 string = "Constante 1"

	// substituindo os valores de uma variavel e alterar a ordem

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(constante1)
}
