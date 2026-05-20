package main

import (
	"fmt"
)

// Um map é uma estrutura de dados que funciona por meio de chave/valor, em que cada chave possui um valor correspondente
// Um map possui a ordem que os itens foram colocados na MAIORIA das vezes, pode acontecer de vir uma ordem aleatória

func Utilizando_map() {
	amigos := map[string]int{
		"Ana":   12345678,
		"Jorge": 87654321,
	}

	fmt.Println(amigos)

	// Para adicionar itens em um map é só colocar dessa forma:
	amigos["Leonardo"] = 18273645

	fmt.Println(amigos)

	// Para verificar se o valor existe ou não é utilizado um conceito chamado "comma ok" que manda um bool de se a chave foi achada ou não
	pessoa, ok := amigos["fantasma"]

	fmt.Println(pessoa, ok)
}

// Para retirar um elemento de um map, usa-se a função delete([nome_do_map], [elemento])
