package main

import (
	"fmt"
)

// Os ponteiros são a localização dos dados dentro da memória
// Para acessar o ponteiro de uma variável é só usar &[variável]

func Ponteiros() {
	x := 10

	fmt.Println(&x)

	// Os ponteiros podem ser salvos dentro de outras variáveis
	y := &x

	// Para saber o conteúdo do ponteiro é só deferir a variável com *[variável]
	fmt.Println(*y)
}
