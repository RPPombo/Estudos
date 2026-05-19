package exercicios

import (
	"fmt"
)

func Exercicio_3() {
	ano_atual := 2026
	ano := 2005

	for ano <= ano_atual {
		fmt.Println(ano)
		ano++
	}
}
