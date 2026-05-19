package exercicios

import (
	"fmt"
)

func Exercicio_8() {
	situacao := 3

	switch {
	case situacao == 1:
		fmt.Println("Chame os bombeiros!")
	case situacao == 2:
		fmt.Println("Chame uma ambulância!")
	case situacao == 3:
		fmt.Println("Chame a polícia!")
	case situacao == 0:
		fmt.Println("Não precisa chamar ninguém")
	default:
		fmt.Println("Situação desconhecida")
	}
}
