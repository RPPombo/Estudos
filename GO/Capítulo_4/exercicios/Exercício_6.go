package exercicios

import (
	"fmt"
)

func Exercicio_6() {
	const (
		_ = iota + 2026
		a1
		a2
		a3
		a4
	)

	proximos_anos := fmt.Sprintf("%d\n%d\n%d\n%d\n", a1, a2, a3, a4)

	fmt.Println(proximos_anos)
}
