package exercicios

import (
	"fmt"
)

func Exercicio_5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d %% 4 = %d\n", i, i%4)
	}
}
