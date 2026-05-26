package main

import (
	"context"
	"fmt"
)

// O context é uma forma de controlar as goroutines criadas
// Ele é capaz de cancelar uma operação com dezenas de goroutines, fazendo com que todas sejam canceladas

func Basico_de_context() {
	ctx := context.Background()

	fmt.Println("Context:", ctx)
	fmt.Println("Context err:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)

	// Adicionando a opção de cancelamento ao contexto
	// A variável cancel, representa a função responsável por cancelar as goroutines relacionadas ao context.Background
	ctx, cancel := context.WithCancel(ctx)

	cancel()

	fmt.Println("Context:", ctx)
	fmt.Println("Context err:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
}
