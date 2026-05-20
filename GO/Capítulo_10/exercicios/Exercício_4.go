package exercicios

import (
	"fmt"
)

func Exercicio_4() {
	anonimo := struct {
		matriz_maluca map[int]string
		lista         []int
	}{
		matriz_maluca: map[int]string{1: "Alguma coisa", 2: "Outra coisa"},
		lista:         []int{1, 2, 3, 4},
	}

	fmt.Println(anonimo)
}
