package exercicios

import (
	"fmt"
)

func Exercicio_6() {
	canal := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			canal <- i
		}
		close(canal)
	}()

	for v := range canal {
		fmt.Println("Recebido:", v)
	}

	fmt.Println("Acabou!")
}
