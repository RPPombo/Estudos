package main

import (
	"fmt"
)

// O tipo booleano é responsável pelas comparações, guardando o resultado de "true" ou "false" nele

func tipo_booleano() {
	var x bool
	fmt.Println(x)

	x = true
	fmt.Println(x)

	x = 10 < 1
	y := 500 == 600
	z := 80 != 7

	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)
}
