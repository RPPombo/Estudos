package exercicios

import (
	"fmt"
	"runtime"
)

func Exercicio_6() {
	os := runtime.GOOS
	arch := runtime.GOARCH

	fmt.Println("Você está rodando esse programa no Sistema Operacional:", os)
	fmt.Println("A arquitetura do seu computador é:", arch)
}
