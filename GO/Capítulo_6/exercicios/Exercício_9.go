package exercicios

import (
	"fmt"
)

func Exercicio_9() {
	var esporte_favorito string = "basquete"

	switch esporte_favorito {
	case "futebol":
		fmt.Println("Fã de Brasileirão")
	case "basquete":
		fmt.Println("Fã de NBA")
	case "fórmula 1":
		fmt.Println("Fã de alta velocidade")
	default:
		fmt.Println("Esporte interessante")
	}
}
