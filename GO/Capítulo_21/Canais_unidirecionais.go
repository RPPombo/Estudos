package main

import (
	"fmt"
)

// É possível especificar se o canal é um canal de apenas leitura(<-chan type) ou apenas escrita(chan<- type)
// É possível fazer com que um canal bidirecional se passe por específicos, mas não o contrário

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	num := <-c

	fmt.Println("Número recebido:", num)
}

func Canais_unidirecionais() {
	canal := make(chan int)

	go send(canal)

	receive(canal)
}
