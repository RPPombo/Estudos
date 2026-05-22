package exercicios

import (
	"fmt"
)

func Funcao_int() int {
	x := 10
	y := 6

	return x * y
}

func Funcao_int_str() (int, string) {
	numero := 99
	frase := "contra germes"

	return numero, frase
}

func Exercicio_1() {
	multiplicacao := Funcao_int()
	porcentagem, usado := Funcao_int_str()

	fmt.Println("O resultado de 6X10 é", multiplicacao)
	println("Sabonete é usado", usado, "e tem uma eficácia de", porcentagem, "%")
}
