package main

import (
	"fmt"
)

// Os canais são maneiras de transmitir dados entre goroutines, criando assim uma sincronização e compartilhamento entre elas
// Para adicionar uma informação em um canal, alguma routine tem que estar esperando e outra enviando, não pode ter apenas uma
// Outra funcionalidade que pode ser adicionado à um canal é a inserção de um buffer, que guarda n informações até serem retiradas

func Introducao_canais() {
	// make(chan type, buffer), o type é o tipo de informação que o canal vai transmitir
	canal := make(chan int)

	go func() {
		canal <- 42
	}()

	fmt.Println(<-canal)

}

// Os canais, assim como em maps, devolvem valores com uma variável de controle (comma ok)
// Os canais podem ser convergidos em outro canal para a saída ser feita somente em um lugar
// Outro conceito importante, é o de divergência, em que para cada valor enviado em um canal, uma goroutine é criada, e após isso todas convergem em um outro canal
