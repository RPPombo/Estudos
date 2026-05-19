package main

import (
	"fmt"
)

// Os operadores lógicos são muito importantes para condições em programação
// Eles comparam valores booleanos para retornar um valor booleano
// Os operadores são AND(&&) e OR(||)

func Operadores_logicos() {
	a := true
	b := true

	if a && b {
		fmt.Println("A e B são verdadeiros")
	} else if a || b {
		fmt.Println("A ou B é verdadeiro")
	}
}
