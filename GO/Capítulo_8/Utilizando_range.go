package main

import (
	"fmt"
)

// O statement range retorna cada valor e indice dos elementos de um slice

func Utilizar_range() {
	slice := []string{"banana", "maçã", "jaca"}

	for indice, valor := range slice {
		fmt.Printf("No índice %d está salvo: %s\n", indice, valor)
	}
}
