package exercicios

import (
	"fmt"
)

func Exercicio_1() {
	type pessoa struct {
		nome              string
		sobrenome         string
		sabores_favoritos []string
	}

	pessoa1 := pessoa{"Amanda", "Silva", []string{"baunilha", "chocolate"}}
	pessoa2 := pessoa{"Jonas", "Pinheiro", []string{"manga", "limão"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.sabores_favoritos {
		fmt.Println(v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.sabores_favoritos {
		fmt.Println(v)
	}
}
