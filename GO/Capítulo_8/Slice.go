package main

import (
	"fmt"
)

// Um slice é uma estrutura de dados que guarda dados de um mesmo tipo,porém com um tamanho ilimitado
// Como a estrutura "slice" é baseada em "array", não é possível acessar um índice que não foi inicializado, mesmo com o slice sendo ilimitado

func Slices() {
	x := []int{1, 2, 3}
	fmt.Println(x)

	// A função append() adiciona elementos no final do slice
	x = append(x, 4, 5, 6)
	fmt.Println(x)
}
