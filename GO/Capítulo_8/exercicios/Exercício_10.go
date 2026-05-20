package exercicios

import (
	"fmt"
)

func Exercicio_10() {
	matriz := map[string][]string{}

	matriz["ferraz_anderson"] = []string{"Natação", "Corrida", "Musculação"}
	matriz["dias_carlos"] = []string{"Jogar Videogame", "Passear", "Dormir"}
	matriz["moraes_eliana"] = []string{"Dançar", "Beber", "Cozinhar"}

	for i, v := range matriz {
		fmt.Println("Pessoa:", i)
		fmt.Println("Hobbies:", v)
	}

	fmt.Printf("\n")

	matriz["gomez_heitor"] = []string{"Tocar", "Fumar", "Jogar Futebol"}

	for i, v := range matriz {
		fmt.Println("Pessoa:", i)
		fmt.Println("Hobbies:", v)
	}

	fmt.Printf("\n")

	delete(matriz, "dias_carlos")

	for i, v := range matriz {
		fmt.Println("Pessoa:", i)
		fmt.Println("Hobbies:", v)
	}
}
