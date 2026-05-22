package exercicios

import (
	"fmt"
)

func executor(funcao func()) {
	funcao()
}

func Exercicio_9() {
	soma := func() {
		x := 10
		y := 70
		fmt.Println("A soma é de", x+y)
	}

	executor(soma)
}
