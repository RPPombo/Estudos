package main

import (
	"fmt"
)

// Os tipos numéricos possuem a característica de delimitar o tamanho deles pela quantidade de bits
// Inteiros: int8, int16, int32, int64 (funcionando para unsigned integer -> uint)
// Floats: float32, float64

// Para caracteres de codificação utf-8, existe o apelido de "rune" ao invés de "int32", para os bytes

// Não é possível usar um int32 como um int, é necessário uma conversão antes

func tipo_numerico() {
	// Por default o compilador deixa como int e float64 as variáveis não declaradas explicitamente
	a := 10
	b := 10.5

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	var overflow uint16
	overflow = 65535
	fmt.Println(overflow)

	overflow++
	fmt.Println(overflow)
}

// Para overflows, quando o número é atribuído direto no código-fonte, ele irá retornar um erro
// Caso o overflow aconteça por meio de alguma operação o número irá loopar e retornar ao número contrário
