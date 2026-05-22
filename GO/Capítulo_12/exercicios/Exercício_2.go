package exercicios

import (
	"fmt"
)

func Soma_ints(numeros ...int) int {
	total := 0
	for _, v := range numeros {
		total += v
	}

	return total
}

func Soma_ints_slice(numeros []int) int {
	total := 0
	for _, v := range numeros {
		total += v
	}

	return total
}

func Exercicio_2() {
	numeros := []int{1, 2, 3, 4, 5}

	soma1 := Soma_ints(numeros...)
	soma2 := Soma_ints_slice(numeros)

	fmt.Println("Somas recebidas:", soma1, soma2)
}
