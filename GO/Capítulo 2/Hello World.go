package main

import (
	"fmt"
)

func main() {
	// O _ serve para descartar valores retornados por funções, nesse caso o número de bytes que está escrito no terminal
	_, erros := fmt.Println("Hello World!")
	fmt.Println(erros)
}
