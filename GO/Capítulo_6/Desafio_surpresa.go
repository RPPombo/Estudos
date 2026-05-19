package main

import (
	"fmt"
)

// É possível fazer a conversão de um int para o caractere correspondente em UTF-8 de forma simples em GO
func Desafio_surpresa() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\t%v\n", i, i, i, string(i))
	}
}
