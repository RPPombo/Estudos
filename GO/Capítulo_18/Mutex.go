package main

import (
	"fmt"
	"runtime"
	"sync"
)

// O mutex funciona de forma que a variável não possa ser alterada enquanto aquela goroutine, não destrancá-la

func Utilizando_mutex() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	total_de_goroutines := 15

	var wg sync.WaitGroup
	wg.Add(total_de_goroutines)

	var mu sync.Mutex

	for i := 0; i < total_de_goroutines; i++ {
		go func() {
			// Trava o espaço de memória utilizado na goroutine
			mu.Lock()
			v := contador
			// Faz com que a goroutine seja parada e volte depois de um tempo
			runtime.Gosched()
			v++
			contador = v

			// Destrava o espaço de memória da goroutine
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

}

// O atomic, por ser a base do mutex, funciona de forma semelhante, porém mais primitiva
// Em que o endereço da variável é travado e para ocorrer alguma mudança lá deve ser incrementado de maneira atômica
// Na prática, o atomic quase nunca é usado
