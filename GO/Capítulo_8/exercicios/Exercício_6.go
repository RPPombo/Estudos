package exercicios

import (
	"fmt"
)

func Exercicio_6() {
	estados := make([]string, 27, 30)

	estados = []string{"Amazonas", "Acre", "Alagoas", "Amapá", "Bahia", "Ceará", "Distrito Federal",
		"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais",
		"Paraíba", "Pará", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Sul",
		"Rio Grande do Norte", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estados), cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}
