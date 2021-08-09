package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	c := cachorro{"Rex", "Doberman", 5}
	fmt.Println(c)

	j, e := json.Marshal(c)

	if e != nil {
		log.Fatal("Ocorreu um erro no Marshal!", e)
	}

	fmt.Println(string(j))

	// fmt.Println(bytes.NewBuffer(j)) // retorno proposto pelo curso

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2, e2 := json.Marshal(c2)

	if e2 != nil {
		log.Fatal("Ocorreu um erro no Marshal!", e)
	}

	fmt.Println(string(cachorro2))
}
