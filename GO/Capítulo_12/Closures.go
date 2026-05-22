package main

import (
	"fmt"
)

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func Fazendo_closure() {
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// O conceito de closure é da criação de um scope diferente, toda vez que uma função é salva em uma variável diferente
	// Isso acontece devido a uma cópia da função ser criada no salvamento dela em uma variável

	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

// Nesse código há dois scopes salvos, o da variável "a" e o da variável "b"
