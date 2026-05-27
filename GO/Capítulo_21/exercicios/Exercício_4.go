package exercicios

import (
	"fmt"
)

func gen4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()

	return c
}

func receive4(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Recebido:", v)
		case _, ok := <-q:
			if ok {
				fmt.Println("Parou de chegar coisa!")
				return
			}
		}
	}
}

func Exercicio_4() {
	q := make(chan int)
	c := gen4(q)

	receive4(c, q)

	fmt.Println("about to exit")
}
