package main

import (
	"fmt"
)

// Em GO é possível fazer o deslocamento de bits de forma simples
func deslocamento_de_bits() {
	x := 2
	y := x >> 1
	z := x << 1

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
}
