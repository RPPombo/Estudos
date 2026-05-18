package exercicios

import (
	"fmt"
)

type tipo_novo int

var x tipo_novo
var y int

func Exercicio_5() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)

	fmt.Println(y, x)
}
