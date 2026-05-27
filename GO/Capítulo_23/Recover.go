package main

import (
	"fmt"
)

func Utilizando_recover() {
	f()
	fmt.Println("Returned normally from f.")
}

// A função recover() faz com que caso ocorra algum pânico, ela retorne um valor não nulo e cancele o pânico a partir daquela função

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
