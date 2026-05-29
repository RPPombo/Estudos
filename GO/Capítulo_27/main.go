package main

import (
	"fmt"
)

func Soma(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}

	return total
}

func main() {
	fmt.Println(Soma())
}
