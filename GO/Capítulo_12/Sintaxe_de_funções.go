package main

// A sintaxe de uma função é composta pelos seguintes itens:
// func (receiver) function_name (parameters) (returns) {code}
// Não necessáriamente precisa-se de parâmetros, retorno e receiver em uma função
// Em Go sempre que se usa uma função, são envidos valores e não referências (endereço de memória)
// Caso um parâmetro possa ser variático, é só colocar "...", após o nome dele, transformando assim ele em uma slice

func soma_subtracao(x int, y int) (int, int) {
	soma := x + y
	subtracao := x - y

	return soma, subtracao
}

// Obs: Em documentações, os valores que entram na função são chamados de parâmetros quando está falando de recebimento
// Quando é sobre envio de valores, esses mesmo valores são chamados de argumentos
