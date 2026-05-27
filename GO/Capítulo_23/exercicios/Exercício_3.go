package exercicios

import (
	"errors"
	"fmt"
	"log"
)

type erro_especial struct {
	erro_dado string
}

func (e erro_especial) Error() string {
	err := fmt.Sprintf("Deu ruim demais saca só: %v", e.erro_dado)
	return err
}

func funcao_esquisita(e error) error {
	err := e.Error()
	return errors.New(err)
}

func Exercicio_3() {
	erro_custom := erro_especial{"criado com sucesso!"}

	erro_recebido := funcao_esquisita(erro_custom)

	log.Println(erro_recebido)
}
