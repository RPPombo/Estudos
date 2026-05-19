package main

import (
	"fmt"
)

// Um array é uma estrutura de dados que comporta dados do mesmo tipo
// O tipo do array é a quantidade de dados e o tipo deles
// É possível acessar um objeto do array explicitando o seu índice

func Arrays() {
	x := [5]int{1, 2, 3, 4, 5}

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// Com a função len([array]) é possível saber o tamanho do array
	fmt.Println(len(x))
}
