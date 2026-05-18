package exercicios

import (
	"fmt"
)

func Exercicio_4() {
	var numero int = 40

	fmt.Printf("%d, %b, %#x\n", numero, numero, numero)

	var numero_deslocado int = numero << 1

	fmt.Printf("%d, %b, %#x\n", numero_deslocado, numero_deslocado, numero_deslocado)
}
