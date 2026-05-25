package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Uma Goroutine é uma forma de criar uma "thread" em GO, não é exatamente uma thread, mas o conceito é parecido
// As funções vão rodar em concorrência(pseudoparalelismo) para tornar o código mais otimizado

var wg sync.WaitGroup

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(200)
	}

	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(200)
	}

	wg.Done()
}

func Utilizando_goroutine() {
	//Verificar a quantidade de Goroutines criadas
	fmt.Println(runtime.NumGoroutine())

	// Essa função adiciona quantos "Dones", deve ser recebido para que o WaitGroup seja liberado
	wg.Add(2)

	// O statement GO, adiciona a função em uma goroutine que fica isolada
	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())

	// A função libera o waitgroup
	wg.Wait()
}

// As Goroutines tem um problema que caso elas utilizem uma mesma variável/espaço de memória, pode gerar uma condição de corrida
// Uma condição de corrida faz com que o código não seja seguro, pois a saída não é garantida
// Para fazer com que não haja esse problema, utiliza-se algumas formas: Mutex, Atomic, Channels
