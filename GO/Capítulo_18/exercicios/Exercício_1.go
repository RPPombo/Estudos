package exercicios

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Exercicio_1() {
	fmt.Println("Essa frase faz parte da goroutine principal!")

	wg.Add(2)

	go contar_ate_10()
	go dar_oi()

	wg.Wait()
}

func contar_ate_10() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func dar_oi() {
	fmt.Println("Oi! Tudo bem?")

	wg.Done()
}
