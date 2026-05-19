package exercicios

import (
	"fmt"
)

func Exercicio_7() {
	x := 10

	if x > 20 {
		fmt.Println("X é maior que 20")
	} else if x > 10 {
		fmt.Println("X é maior que 10 e menor que 20")
	} else {
		fmt.Println("X é menor ou igual à 10")
	}
}
