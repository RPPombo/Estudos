package exercicios

import (
	"fmt"
)

func Exercicio_5() {
	c := make(chan int, 1)

	c <- 50

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
