package exercicios

import (
	"fmt"
)

func Exercicio_7() {
	ss := [][]string{
		[]string{"Paula", "Gustavo", "Marcos"},
		[]string{"Carvalho", "Jesus", "Coutinho"},
		[]string{"Cantar", "Correr", "Nadar"},
	}

	for _, v := range ss {
		fmt.Println(v)
	}
}
