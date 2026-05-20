package exercicios

import (
	"fmt"
)

func Exercicio_2() {
	type pessoa struct {
		nome              string
		sobrenome         string
		sabores_favoritos []string
	}

	matriz := map[string]pessoa{}

	pessoa1 := pessoa{"Amanda", "Silva", []string{"baunilha", "chocolate"}}
	pessoa2 := pessoa{"Jonas", "Pinheiro", []string{"manga", "limão"}}

	matriz[pessoa1.sobrenome] = pessoa1
	matriz[pessoa2.sobrenome] = pessoa2

	for _, valor := range matriz {
		fmt.Println(valor.nome, valor.sobrenome)

		for _, v := range valor.sabores_favoritos {
			fmt.Println(v)
		}
	}

}
