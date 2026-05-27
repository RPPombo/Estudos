package main

import (
	"fmt"
	"log"
)

// A maneira mais fácil de criar um erro customizado é com o package errors, pois ele transforma strings em tipo error
// Para criar um erro a partir de uma string é só usar errors.New()
// Por ser um tipo, assim como qualquer outro, o erro customizado pode ser guardado em uma váriavel/constante
// Caso queira fazer uma mensagem de erro monstrando, por exemplo a variável que causou o problema pode usar fmt.Errorf()

// Outra opção além de todas as outras citadas é criar uma estrutura com o método Error()

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func Erro_custom() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}*/

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}
