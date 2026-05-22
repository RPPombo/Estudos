package exercicios

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func Mude_me(ponteiro *pessoa) {
	fmt.Println("Mudei o seu nome")
	(*ponteiro).nome = "Antonio"
	fmt.Println("Seu nome é", (*ponteiro).nome, "agora")
}

func Exercicio_2() {
	alguem := pessoa{"Carlos", "da Silva", 31}

	Mude_me(&alguem)
}
