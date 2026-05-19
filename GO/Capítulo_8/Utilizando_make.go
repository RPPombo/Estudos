package main

import (
	"fmt"
)

// A função make() serve para criar um slice com n índices e reservar espaço para x de capacidade
// Na prática isso faz com que o custo computacional de ficar destruindo e criando arrays toda vez que o slice aumentar diminua
// Para adicionar elementos na slice é só usar o append() normalmente

func Utilizar_make() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 6)
	fmt.Println(slice)
}

// Caso o len se torne maior que o cap, automaticamente o cap vai ser dobrado para manter o funcionamento da função make()
