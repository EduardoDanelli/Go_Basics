package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	c1 := cachorro{"Tobi", "Rusky", 2}
	fmt.Println(c1)

	caoJSON, erro := json.Marshal(c1)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(caoJSON)
	fmt.Println(bytes.NewBuffer(caoJSON))

	c2 := map[string]string{
		"nome": "Cão",
		"raca": "Pitbul",
	}

	cao2JSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cao2JSON))
}
