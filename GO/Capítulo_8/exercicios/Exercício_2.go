package exercicios

import (
	"fmt"
)

func Exercicio_2() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", slice)
}
