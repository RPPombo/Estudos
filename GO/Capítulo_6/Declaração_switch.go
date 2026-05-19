package main

import (
	"fmt"
)

// O switch serve para criar vários casos diferentes relacionados à um statement ou uma variável
// Por padrão ele compara o case com true, se o switch não for especificado
// O statement falltrough faz com que caso o seja verdadeiro um case, de baixo será também
// Um case pode ter várias formas de ser ativo, para isso basta colocar as formas após a vírgula
// O default, faz com que caso não ocorra nenhuma das opções anteriores, acontecerá essa
func Declaracao_switch() {
	pessoa := "Thomas"
	switch pessoa {
	case "Ana":
		fmt.Println("Financeiro ativo")
		fallthrough
	case "Ronaldo", "Thomas":
		fmt.Println("Produção ativa")
		fallthrough
	case "Marcos", "Pedro":
		fmt.Println("RH ativo")
	default:
		fmt.Println("Ninguém veio hoje!")
	}
}

// Assim como no for, o switch pode ter uma inicialização
