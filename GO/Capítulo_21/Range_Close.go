package main

import (
	"fmt"
)

func Fechando_canais() {
	canal := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 20; i++ {
			c <- i
		}
		// Caso o canal não seja fechado e a goroutine acabe, o programa irá dar um fatal error
		close(c)
	}(canal)

	// O range, quando usado em canais, fica constantemente escutando o canal, até ele ser fechado
	for v := range canal {
		fmt.Println(v)
	}
}
