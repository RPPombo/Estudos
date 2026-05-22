package exercicios

import (
	"fmt"
)

type pessoa_exercicio struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa_exercicio) demonstrar_infos() {
	fmt.Println("Meu nome é", p.nome, p.sobrenome)
	fmt.Println("Eu tenho", p.idade, "anos")
}

func Exercicio_4() {
	henrique := pessoa_exercicio{"Henrique", "Cabral", 70}

	henrique.demonstrar_infos()
}
