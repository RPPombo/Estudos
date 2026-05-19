package main

import (
	"fmt"
)

var x interface{}

// Para verificar o tipo de alguma variável é só utilizar: [variável].(type)
func verificar_tipo() {
	x = true

	switch x.(type) {
	case int:
		fmt.Println("Tipo int")
	case bool:
		fmt.Println("Tipo bool")
	case float64:
		fmt.Println("Tipo float64")
	case string:
		fmt.Println("Tipo string")
	default:
		fmt.Println("Não é um dos tipos principais!")
	}
}
