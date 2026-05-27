package main

import (
	"fmt"
	"os"
)

// Há algumas formas de mostrar os erros quando eles acontecem
// O fmt.Prinln() é a forma mais simples de mostrar o erro, em que ele simplesmente escreve o erro no stdout
// O log.Println() escreve o erro ocorrido com a data e horário no "standard error" (pode ser configurado com log.SetOutput())
// o log.Fatalln() escreve o erro com um println e mata o programa instantaneamente com Exit(1)
// O log.Panic() escreve um println com o erro, para a execução da função que chamou, roda o defer tanto da função atual quanto da principal e fecha o programa com Exit(2)
// O panic() apenas retorna o panic que aconteceu no programa

func Mostrando_erros() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

// Obs: as funções relacionadas com panic podem ter uma forma de recuperar a execução do programa com o recover()
