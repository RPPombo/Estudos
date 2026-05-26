package exercicios

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println("A pessoa", (*p).nome, "está falando!")
}

type humanos interface {
	falar()
}

func dizer_alguma_coisa(h humanos) {
	h.falar()
}

func Exercicio_2() {
	paulo := pessoa{"Paulo", 32}

	dizer_alguma_coisa(&paulo)
}
