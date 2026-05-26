package exercicios

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg4 sync.WaitGroup

func Exercicio_5() {
	wg4.Add(100)

	var contador atomic.Int32
	contador.Store(0)

	for i := 0; i < 100; i++ {
		go func() {
			contador.Add(1)
			runtime.Gosched()
			fmt.Println("Contador:", contador.Load())
			wg4.Done()
		}()
	}

	wg4.Wait()
}
