package exercicios

import (
	"fmt"
)

func Exercicio_3() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, v := range slice[:3] {
		fmt.Println(v)
	}
	fmt.Printf("\n")

	for _, v := range slice[4:] {
		fmt.Println(v)
	}
	fmt.Printf("\n")

	for _, v := range slice[1:7] {
		fmt.Println(v)
	}
	fmt.Printf("\n")

	for _, v := range slice[2 : len(slice)-1] {
		fmt.Println(v)
	}

}
