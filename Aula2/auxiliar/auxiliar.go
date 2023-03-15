package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do Pacote Auxiliar")
	escrever2()
}

// como está usando o mesmo pacote você pode usar a outra função aqui dentro, porém no main não tem como