package exercicios

import (
	"fmt"
)

func funcao_criadora() func() {
	return func() {
		fmt.Println("Eu fui criada por outra função")
	}
}

func Exercicio_8() {
	a := funcao_criadora()

	a()
}
