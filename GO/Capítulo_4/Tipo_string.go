package main

import (
	"fmt"
)

func tipo_string() {
	frase := "Qualquer coisa"

	fmt.Printf("%v, %T\n", frase, frase)

	// Para o uso de strings de forma literal, deve-se utilizar o acento grave (``)
	frase_torta := `Qualquer 
	coisa
				bem torta`

	fmt.Printf("%v, %T\n", frase_torta, frase_torta)

	// Para converter strings em slice de bytes faz da seguinte forma:
	frase_bytes := []byte(frase)

	fmt.Printf("%v, %T\n", frase_bytes, frase_bytes)
}

// É possível passar por caracteres e bytes de uma string(um por um)
// Utilizando um for loop baseado em range, o retorno será o caractere em UTF-8
// Utilizando um for loop básico condicional, o retorno será por byte
// Sendo assim, caracteres fora do código ASCII serão perdidos utilizando a segunda forma
