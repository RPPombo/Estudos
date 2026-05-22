package main

import (
	"fmt"
)

// O conceito de recursividade é uma função chamar ela mesma até um ponto de parada

func fatorial(x int) int {
	if x != 1 {
		return (x * fatorial(x-1))
	} else {
		return 1
	}
}

// Os casos mais simples de entender são o cálculo de fatoriais e dos números de fibonacci

func Utilizando_recursividade() {
	mult := fatorial(4)

	fmt.Println(mult)
}

// Na maioria das vezes é mais aconselhável usar loops, por gastar menos memória e ter menos chances de dar problemas
