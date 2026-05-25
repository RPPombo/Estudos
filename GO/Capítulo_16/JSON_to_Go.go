package main

import (
	"encoding/json"
	"fmt"
)

// Para ler um JSON e salvá-lo em um struct é necessário especificar no struct que os campos vão ser preenchidos por variáveis vindas dele
// Essas tags também são úteis para especificar os campos quando a struct for exportada como JSON

type informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func Ler_JSON() {
	json_recebido := []byte(`[{"Nome":"João","Sobrenome":"Ferreira","Idade":32,"Profissao":"arquiteto","Contabancaria":5000.9}, 
	{"Nome":"Paulo","Sobrenome":"da Silva","Idade":20,"Profissao":"programador","Contabancaria":7000.65}]`)

	var json_lido []informacoes

	err := json.Unmarshal(json_recebido, &json_lido)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(json_lido)
}
