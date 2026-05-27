package exercicios

import (
	"fmt"
)

func gen(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- i
	}
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("Recebido:", v)
	}
}

func Exercicio_3() {
	canal := make(chan int)

	go receive(canal)
	gen(canal)

	fmt.Println("Acabou!")
}
