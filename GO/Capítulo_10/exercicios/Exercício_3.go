package exercicios

import (
	"fmt"
)

func Exercicio_3() {
	type veiculo struct {
		portas int
		cor    string
	}

	type caminhonete struct {
		veiculo
		quatro_rodas bool
	}

	type sedan struct {
		veiculo
		modelo_de_luxo bool
	}

	strada := caminhonete{veiculo{2, "prata"}, true}
	corolla := sedan{veiculo{4, "preto"}, false}

	fmt.Println("Quantidade de portas:", strada.veiculo.portas)
	fmt.Println("Cor:", strada.veiculo.cor)
	fmt.Println("Quatro rodas:", strada.quatro_rodas)

	fmt.Println("Quantidade de portas:", corolla.veiculo.portas)
	fmt.Println("Cor:", corolla.veiculo.cor)
	fmt.Println("Modelo de luxo", corolla.modelo_de_luxo)
}
