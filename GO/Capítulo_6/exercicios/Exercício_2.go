package exercicios

import (
	"fmt"
)

func Exercicio_2() {
	for letra := 65; letra <= 90; letra++ {
		fmt.Println(letra)
		repeticao := 0
		for repeticao < 3 {
			fmt.Printf("%#U\n", letra)
			repeticao++
		}
	}
}
