package main

import (
	"fmt"
)

// Para declarar variáveis em package level scope é utilizado a palavra "var"
var global = 30.0

func hello_world() {
	// O := (gother) serve para autodeclarar uma variável com sua tipagem e atribuir um valor à ela
	// Para ele ser usado, ele precisa declarar pelo menos uma variável e só funcionará dentro de codeblocks
	x := 12
	y := "Hello World"

	// O _ serve para descartar valores retornados por funções, nesse caso o número de bytes que está escrito no terminal
	_, erros := fmt.Printf("x: %d %T\n", x, x)
	fmt.Printf("y: %s %T\n", y, y)
	fmt.Printf("z: %f %T\n", global, global)
	fmt.Println(erros)

}
