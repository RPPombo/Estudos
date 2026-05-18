package main

import (
	"fmt"
)

// Para criar tipos, deve-se colocar o nome do tipo e em qual existente ele se baseia

type meu_tipo int

var tipo_diferente meu_tipo

// Os tipos em GO são estáticos, ou seja não é possível atribuir valores de tipos diferentes das variáveis
// O compilador já deduz automaticamente o tipo da variável na hora em que é atribuído um valor, mas pode e deve ser declarado caso necessário

var numero int
var valor_zero float32

// Caso a variável tenha sido declarada, porém não tenha recebido nenhuma atribuição, ela terá um valor "zero" nela

func tipo() {
	// A atribuição só pode ser feita dentro de codebloccks
	numero = 10

	fmt.Printf("%d\n", numero)

	fmt.Printf("%f\n", valor_zero)

	fmt.Printf("%T\n", tipo_diferente)
}

// Em GO apenas tem a conversão de tipos, para ela ocorrer basta colocar [tipo_desejado](vaeriável) e salvar em uma nova variável
