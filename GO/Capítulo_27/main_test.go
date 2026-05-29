package main

import (
	"fmt"
	"testing"
)

// Para fazer testes de um arquivo, deve-se escrever o nome do arquivo com a terminação "_test.go"
// O arquivo de testes deve estar no mesmo package que o arquivo a ser testado
// Para testar funções o nome delas deve ser transformado em "Test[nome da função]" e deve-se enviar um ponteiro de testing.T

// Para fazer vários testes em sequência, cria-se uma struct test{data type, answer type}
type test struct {
	data   []int
	answer int
}

func Test_Soma(t *testing.T) {
	x := Soma(3, 2, 1)
	resultado := 6

	if x != resultado {
		t.Errorf("Expexted: %v. Got: %v", resultado, x)
	}
}

// Função com vários testes sequenciais
func Test_Soma2(t *testing.T) {
	tests := []test{
		{[]int{12, 25, 5}, 42},
		{[]int{80, 100, 10}, 190},
		{[]int{75, 25, 200}, 300},
	}

	for _, v := range tests {
		x := Soma(v.data...)
		resultado := v.answer

		if x != resultado {
			t.Errorf("Expexted: %v. Got: %v", resultado, x)
		}
	}

}

// Para fins de documentação, o nome de exemplos de testes, deve ser "Example[nome da função]"
// Outro fator importante é colocar um comentário "Output: saída"

func ExampleSoma() {
	fmt.Println(Soma(1, 2, 3))
	//Output: 6
}

// O comando de teste no terminal é "go test" ou "go test -v"

// Benchmarks são responsáveis por retornar parâmetros de qualidade, para verificar se a função está adequada para o projeto

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(79, 80, 54)
	}
}

// Para rodar um Benchmark, utiliza-se "go test -bench" ou "go test -bench [nome da função]"

// Outro parâmetro bom de se analisar é o de cobertura de testes do código, em média 70-80% do código é coberto
// Comando: "go test -cover"
// Para salvar em um arquivo, utilize:"go test -coverprofile [nome do arquivo]"
