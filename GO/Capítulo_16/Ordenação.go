package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordem_por_nome []carro

func (x ordem_por_nome) Len() int {
	return len(x)
}

// A principal função para a ordenação é essa
func (x ordem_por_nome) Less(i int, j int) bool {
	return x[i].nome < x[j].nome
}

func (x ordem_por_nome) Swap(i int, j int) {
	var pivo carro

	// Maneira de faer sem usar pivô: x[i], x[j] = x[j], x[i]

	pivo = x[i]
	x[i] = x[j]
	x[j] = pivo
}

// O package sort é o pacote responsável por ter funções relacionadas à ordenação

func Ordenar_slices() {
	lista := []string{"Phineas", "Ferb", "Candace", "Isabella", "Perry", "Doofensmirtz"}

	fmt.Println(lista)

	// A partir da versão 1.22 do GO, existe a função slices.Sort() que ordena qualquer tipo de slice
	sort.Strings(lista)

	fmt.Println(lista)
	fmt.Println()

	// Para criar uma ordenação própria, deve-se utilizar a interface sort.Interface e implemntar os métodos necessários para a função sort.Sort()

	lista2 := []carro{
		{"Mustang", 400, 10},
		{"Camaro", 450, 8},
		{"Peugeot", 2, 4},
		{"Strada", 500, 9},
	}

	fmt.Println(lista2)

	sort.Sort(ordem_por_nome(lista2))

	fmt.Println(lista2)
}
