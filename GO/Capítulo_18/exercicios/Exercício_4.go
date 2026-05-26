package exercicios

import (
	"fmt"
	"runtime"
	"sync"
)

var wg3 sync.WaitGroup

func Exercicio_4() {
	wg3.Add(100)

	var mu sync.Mutex
	contador := 0

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			x := contador
			runtime.Gosched()
			x++
			contador = x
			fmt.Println("Contador:", x)
			mu.Unlock()
			wg3.Done()
		}()
	}

	wg3.Wait()
}
