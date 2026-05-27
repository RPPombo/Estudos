package exercicios

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func Exercicio_2() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return bs, fmt.Errorf("Deu esse erro aqui: %v", err)
	}

	return bs, nil
}
