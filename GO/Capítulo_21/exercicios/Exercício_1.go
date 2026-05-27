package exercicios

import (
	"fmt"
)

func Exercicio_1() {
	c := make(chan int, 1)

	/*
		go func() {
			c <- 42
		}()
	*/

	c <- 42

	fmt.Println(<-c)
}
