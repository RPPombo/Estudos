package exercicios

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenar_por_sobrenome []user

func (x ordenar_por_sobrenome) Len() int               { return len(x) }
func (x ordenar_por_sobrenome) Less(i int, j int) bool { return x[i].Last < x[j].Last }
func (x ordenar_por_sobrenome) Swap(i int, j int)      { x[i], x[j] = x[j], x[i] }

type ordenar_por_idade []user

func (x ordenar_por_idade) Len() int               { return len(x) }
func (x ordenar_por_idade) Less(i int, j int) bool { return x[i].Age < x[j].Age }
func (x ordenar_por_idade) Swap(i int, j int)      { x[i], x[j] = x[j], x[i] }

type ordenar_falas []string

func (x ordenar_falas) Len() int               { return len(x) }
func (x ordenar_falas) Less(i int, j int) bool { return x[i] < x[j] }
func (x ordenar_falas) Swap(i int, j int)      { x[i], x[j] = x[j], x[i] }

func Exercicio_5() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(ordenar_por_sobrenome(users))

	fmt.Println(users)

	sort.Sort(ordenar_por_idade(users))

	fmt.Println(users)

	for i := range users {
		sort.Sort(ordenar_falas(users[i].Sayings))
		fmt.Println(users[i].First, users[i].Last)
		for _, v := range users[i].Sayings {
			fmt.Println(v)
		}
	}
}
