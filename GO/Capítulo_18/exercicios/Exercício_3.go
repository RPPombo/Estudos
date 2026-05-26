package exercicios

import (
	"fmt"
	"runtime"
	"sync"
)

var wg2 sync.WaitGroup

func Exercicio_3() {
	wg2.Add(100)

	contador := 0

	for i := 0; i < 100; i++ {
		go func() {
			x := contador
			runtime.Gosched()
			x++
			contador = x
			fmt.Println("Contador:", x)
			wg2.Done()
		}()
	}

	wg2.Wait()
}
