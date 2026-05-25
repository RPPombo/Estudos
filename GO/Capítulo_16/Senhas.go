package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Para importar packages de fora, utilizar o comando "go get [loacalização do package]"

func Criar_hash() {
	senha := "Qualquer coisa"

	// Essa função faz com que uma []byte se transforme em uma hash
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), 4)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(hash)

	entrada := "Qualquer coisa 2"

	// Compara a hash salva com a senha inserida, caso a saída seja nil, significa que está certa
	saida := bcrypt.CompareHashAndPassword(hash, []byte(entrada))

	fmt.Println(saida)
}
