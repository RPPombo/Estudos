package main

import (
	"fmt"
)

// O defer é um statement utilizado para fazer com que ações fiquem para última hora, ou seja, só vão executar depois de todas as outras ações
// Caso haja vários defer, eles utilizaram a lógica de uma pilha, em que o primeiro a entrar, vai ser último a sair
// Em funções que possuem um retorno, o defer vai executar uma linha antes do return

func Utilizando_defer() {
	defer fmt.Println("Print com defer")
	fmt.Println("Print sem defer")
}
