package exercicios

import (
	"fmt"
)

func Exercicio_8() {
	matriz := map[string][]string{}

	matriz["ferraz_anderson"] = []string{"Natação", "Corrida", "Musculação"}
	matriz["dias_carlos"] = []string{"Jogar Videogame", "Passear", "Dormir"}
	matriz["moraes_eliana"] = []string{"Dançar", "Beber", "Cozinhar"}

	for i, v := range matriz {
		fmt.Println("Pessoa:", i)
		fmt.Println("Hobbies:", v)
	}
}
