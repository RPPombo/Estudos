package exercicios

import (
	"fmt"
)

func contador() func() {
	x := 0
	return func() {
		fmt.Println(x)
		x++
	}
}

func Exercicio_10() {
	c := contador()

	c()
	c()
	c()

	c2 := contador()

	c2()
	c2()
	c2()
}
