package exercicios

import (
	"fmt"
)

func Exercicio_7() {
	x := 40

	y := func(num int) {
		quad := x * x
		fmt.Println("O quadrado de", x, "é", quad)
	}

	y(x)
}
