package main

import "fmt"

func Utilizando_select() {
	a := make(chan int)
	b := make(chan int)

	x := 500

	go func(c chan<- int, qtd int) {
		for i := 0; i < qtd; i++ {
			c <- i
		}
	}(a, x/2)

	go func(c chan<- int, qtd int) {
		for i := 0; i < qtd; i++ {
			c <- i
		}
	}(b, x/2)

	for j := 0; j < x; j++ {
		// O select funciona de forma semelhante ao switch, porém com a ideia voltada ao recebimento de vários canais simultaneamente
		select {
		case v := <-a:
			fmt.Println("Canal a mandou:", v)
		case v := <-b:
			fmt.Println("Canal b mandou:", v)
		}
	}
}

// Uma forma interessante de se usar select é criar um canal de controle, em que caso chegue algo nele a função irá fazer outra ação
