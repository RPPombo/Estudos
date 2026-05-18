package exercicios

import (
	"fmt"
)

func Exercicio_2() {
	x := 10
	y := 20

	igualdade := x == y
	menor_igual := x <= y
	menor := x < y
	maior_igual := x >= y
	maior := x > y
	diferente := x != y

	fmt.Println(igualdade)
	fmt.Println(menor_igual)
	fmt.Println(menor)
	fmt.Println(maior_igual)
	fmt.Println(maior)
	fmt.Println(diferente)
}
