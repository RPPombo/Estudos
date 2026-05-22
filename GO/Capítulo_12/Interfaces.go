package main

import (
	"fmt"
)

type dentista struct {
	pessoa
	espec   string
	salario float64
}

type engenheiro struct {
	pessoa
	tipo    string
	salario float64
}

func (p dentista) oibomdia() {
	fmt.Println(p.nome, "está dando bom dia!")
	fmt.Println("Ele é um dentista espcializado em:", p.espec)
}

func (p engenheiro) oibomdia() {
	fmt.Println(p.nome, "está dando bom dia!")
	fmt.Println("Ele é um engenheiro", p.tipo)
}

// A interface é um tipo que se passa por outros tipos, fazendo asssim com funções possam receber vários tipos diferentes, porém com os mesmos argumentos
// A única coisa que a interface precisa ter é os métodos que os tipos que ela pode receber possuem

type gente interface {
	oibomdia()
}

// Nesse caso, o métoddo que ambos os tipos possuem é oibomdia()
// Caso um tipo possuisse um método exclusivo, o método deve ser declarado igualmente, porém somente aquele tipo irá conseguir usar
// No caso de métodos exlcusivos é comum usar switch([variavel].(type)) para que não haja problemas

func serhumano(g gente) {
	g.oibomdia()
}

func Utilizando_interfaces() {
	engenheiro_maluco := engenheiro{pessoa{"Antonio", "Doidão", 999}, "da computação", 4550.87}
	dentista_safado := dentista{pessoa{"Leandro", "Cansado", 800}, "arrancar dentes", 6700.02}

	serhumano(engenheiro_maluco)
	serhumano(dentista_safado)

	engenheiro_maluco.oibomdia()
	dentista_safado.oibomdia()

}
