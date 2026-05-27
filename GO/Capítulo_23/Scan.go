package main

import (
	"fmt"
)

// O tratamento de erros é uma prática muito importante para qualquer linguagem de programção
// Em GO, os desenvolvedores decidiram não criar a estrutura try-catch-finally para forçar devs a ter um código mais organizado
// O trataento de erros geralmente é feito a partir do retorno de uma variável do tipo error em ua função
// É possível criar os próprios erros a partir da interface error, com o método Error()

func Erros_no_scan() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

}
