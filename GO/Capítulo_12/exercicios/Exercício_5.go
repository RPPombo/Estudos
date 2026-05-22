package exercicios

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return c.raio * math.Pi * 2
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func Exercicio_5() {
	quadradao := quadrado{5}
	circulao := circulo{2.5}

	fmt.Println(info(quadradao))
	fmt.Println(info(circulao))
}
