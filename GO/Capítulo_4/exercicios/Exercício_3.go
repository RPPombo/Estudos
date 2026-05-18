package exercicios

import (
	"fmt"
)

func Exercicio_3() {
	const tipada1 int = 10
	const tipada2 float64 = 32

	const ntipada1 = 90
	const ntipada2 = 54.0007

	fmt.Printf("%v => %T\n", tipada1, tipada1)
	fmt.Printf("%v => %T\n", tipada2, tipada2)
	fmt.Printf("%v => %T\n", ntipada1, ntipada1)
	fmt.Printf("%v => %T\n", ntipada2, ntipada2)
}
