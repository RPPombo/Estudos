package exercicios

import (
	"fmt"
)

func Exercicio_3() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range lista {
		if v%2 == 0 {
			fmt.Println(v)
		} else {
			defer fmt.Println(v)
		}
	}
}
