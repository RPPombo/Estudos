package main

import (
	"fmt"
)

// Funções anônimas são funções descartáveis que rodam uma vez, a maior utilidade dela é em go routines

func Funcoes_anonimas() {
	x := 550

	func(x int) {
		fmt.Println(x, "vezes 80 é:")
		fmt.Println(x * 80)
	}(x)

	// Uma função anônima pode ser salva como uma expressão, atribuindo ela à uma variável
	y := func(x int) {
		fmt.Println(x, "vezes 80 é:")
		fmt.Println(x * 80)
	}

	y(x)
}

// Como funções podem ser tratadas como expressões e salvas em variáveis, isso faz com que seja possível retornar funções em funções
// Outra utilidade nisso é o conceito de callback, em que uma função é passada como parâmetro em outra função
