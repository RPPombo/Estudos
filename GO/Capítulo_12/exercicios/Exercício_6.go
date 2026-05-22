package exercicios

import (
	"fmt"
)

func Exercicio_6() {
	x := 30

	func(num int) {
		quad := x * x
		fmt.Println("O quadrado de", x, "é", quad)
	}(x)
}
