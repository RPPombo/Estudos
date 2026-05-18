package main

import (
	"fmt"
	"runtime"
)

// A biblioteca runtime posssui formas de ver informações sobre a máquina que está rodando o programa

func verificar_computador() {
	fmt.Println(runtime.GOOS)   // Retorna o OS
	fmt.Println(runtime.GOARCH) // Retorna a arquitetura
}
