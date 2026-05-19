package main

import (
	"fmt"
)

// O ; em Go funciona com delimitador de statement, em que ele mostra aonde um statement acabou
// Isso faz com que seja possível colocar várias instruções em uma mesma linha

func Loops() {
	// Para fazer um loop for deve-se colocar uma inicialização, um limite e um incremento
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Um loop hierárquico é aquele em que vários loops estão encadeados
	for horas := 0; horas < 24; horas++ {
		for minutos := 0; minutos < 60; minutos++ {
			for segundos := 0; segundos < 60; segundos++ {
				fmt.Printf("%d:%d:%d  ", horas, minutos, segundos)
			}
		}
		fmt.Println("")
	}

	// Em GO, a keyword while não existe, porém é possível fazer um loop while utilizando o for
	condicao := 0

	for condicao < 10 {
		fmt.Println("Condição Verdadeira")
		condicao++
	}
}

// Para quebrar qualquer loop, basta usar o statement "break"
// Para passar direto para a próxima volta do loop, é só usar o statement "continue"
