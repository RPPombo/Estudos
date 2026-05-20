package main

import (
	"fmt"
)

// É possível adicionar structs dentro de outras structs, criando assim tipos de dados mais complexos

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	empregado pessoa
	titulo    string
	salario   float64
}

func Embutindo_structs() {
	pessoa1 := pessoa{"Ântonio", 21}
	pessoa2 := pessoa{"Juliana", 40}

	profissional1 := profissional{pessoa1, "Aprendiz", 1450.50}
	profissional2 := profissional{pessoa2, "Líde de Projeto", 5700.90}

	fmt.Println(profissional1)
	fmt.Println(profissional2)

}

// Em structs embutidos, caso não haja conflito entre o nome dos campos das structs não é necessário especificar que um campo está dentro da outra struct
// Nesse caso, pode-se usar profissional.nome, e irá funcionar
