package main

import (
	"fmt"
)

func constantes() {
	// Uma constante não pode ser mudada em nenhum momento da execução do programa
	// Uma constante só tem sua tipagem definida, quando ela é utilizada no código
	const constante1 = 10

	const (
		constante2 = 54.0007
		constante3 = 70
	)

	var a float64

	a = constante1
	fmt.Printf("%v, %T\n", a, a)

	a = constante2
	fmt.Printf("%v, %T\n", a, a)

	a = constante3
	fmt.Printf("%v, %T\n", a, a)

	// Outra forma de usar constantes, são com IOTAs
	// IOTAs são números inteiros com a tipagem ainda não especifícada(int ou float) que crescem de forma sucessiva

	// É possível colocar fórmulas para o número de cada iota e descartar números com _
	const (
		i1 = iota * 10
		i2
		_
		i3
	)

	println(i1, i2, i3)
}
