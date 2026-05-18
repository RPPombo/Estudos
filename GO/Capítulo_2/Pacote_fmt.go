package main

import (
	"fmt"
)

func pacote(x string, y string) {
	// O Sprint retorna as variáveis juntas em uma string
	retorno1 := fmt.Sprint(x, y)
	// O Sprintln retorna as variáveis em uma string, separadas por um espaço entre elas
	retorno2 := fmt.Sprintln(x, y)
	// O Sprintf retorna as variáveis em uma string, baseado na formatação recebida
	retorno3 := fmt.Sprintf("%s alguma coisa %s", x, y)

	// O Print apenas escreve no terminal
	fmt.Print(retorno1)
	// O Println escreve no terminal e adiciona uma linha nova
	fmt.Println(retorno2)
	// O Printf escreve no terminal com a formatação descrita
	fmt.Printf("%s\n", retorno3)

	// Os "Fprints" seguem a mesma lógica descrita acima, com o detalhe de que é usado para write interfaces
}
