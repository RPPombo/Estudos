package main

import (
	"fmt"
)

// Para pegar partes de slices, basta colocar o índice inicial ":" índice final+1
// Uma maneira de excluir valores de uma slice é utilizando a função append(), junto com as fatias que devem ficar
// É importante saber que quando uma slice vai ser adicionada à outra com função append(), deve-se usar o enumerate(...) para "desempacotar" os valores

func Fatiar_slices() {
	sabores := []string{"mussarela", "napolitana", "marguerita", "lombo", "pepperoni"}

	fatias := sabores[0:2]
	fmt.Println(fatias)

	sabores = append(sabores[:3], sabores[4:]...)
	fmt.Println(sabores)
}
