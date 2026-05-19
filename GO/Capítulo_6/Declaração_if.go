package main

import (
	"fmt"
)

// O if serve para condicionar ações dentro do código
// É muitas vezes usado dentro de loops para condicionar "breaks" e "continues"

func Declaracao_if() {
	x := 10

	if !(x > 100) {
		fmt.Println("X é menor que 100!")
	}

	// É possível adicionar uma inicialização antes do if rodar
	// Para adicionar várias opções em condições deve-se encadeá-las com "else if" e "else"
	if y := 120; y != x {
		fmt.Println("Y é diferente de X!")
	} else {
		fmt.Println("Y é igual a X!")
	}
}

// A condição do if pode ser escrita denro ou fora de parênteses por se tratar de uma comparação matemática
// A ! é o operador lógico "not", que inverte o booleano
