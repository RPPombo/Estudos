package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Em GO tudo que está em letra miúscula pode ser exportado e fica visível para outros packages
// É importante que quando estiver usando JSON, os campos da struct estejam com letra maiúscula, senão não irá funcionar

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func Criar_JSON() {
	joao := pessoa{"João", "Ferreira", 32, "arquiteto", 5000.90}
	paulo := pessoa{"Paulo", "da Silva", 20, "programador", 7000.65}

	j, err := json.Marshal(joao)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	p, err := json.Marshal(paulo)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(string(j))
	fmt.Println(string(p))

	// Outra maneira de utilizar criar um JSON é com um encoder, que tira a necessidade de salvar em uma variável intermediária

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(paulo)

}
