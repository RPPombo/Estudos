package exercicios

import (
	"fmt"
)

func Exercicio_1() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, v := range array {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", array)
}
