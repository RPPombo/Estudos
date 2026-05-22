package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// Um método é uma função ligada diretamente a um tipo
// Apenas se houver uma variável do tipo especificado no receiver da função ela poderá ser utilizada

func (p pessoa) oi() {
	fmt.Println(p.nome, "diz oi!")
}

func Utilizando_metodos() {
	mauricio := pessoa{"Maurício", "Silva", 40}

	// Utilização de um método
	mauricio.oi()
}
