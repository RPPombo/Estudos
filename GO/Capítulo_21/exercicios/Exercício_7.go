package exercicios

import (
	"fmt"
	"sync"
)

func Exercicio_7() {
	var wg sync.WaitGroup
	canal := make(chan int)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
			wg.Done()
		}()
	}

	for contador := 0; contador < 100; contador++ {
		fmt.Println("Recebdido:", <-canal)
	}

	wg.Wait()
}
