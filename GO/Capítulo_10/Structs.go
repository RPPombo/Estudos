package main

import (
	"fmt"
)

// Uma struct é uma estrutura de dados que guarda outros tipos de dados iguais ou não

type cliente struct {
	nome    string
	mesa    int
	fumante bool
}

func Utilizando_structs() {
	cliente1 := cliente{
		nome:    "João",
		mesa:    10,
		fumante: false,
	}

	cliente2 := cliente{
		nome:    "Paula",
		mesa:    11,
		fumante: true,
	}

	fmt.Printf("%v\n%v\n", cliente1, cliente2)

	// Para pegar somente um campo específico do struct, basta colocar: [váriavel].[campo_desejado]
	fmt.Println(cliente1.nome)

	// É possível criar structs sem nessáriamente criar um tipo novo para elas, são chamados de "anonymous structs"
	anonimo := struct {
		nome  string
		idade int
	}{
		nome:  "Segredo",
		idade: 1000000,
	}

	fmt.Println(anonimo)
}
