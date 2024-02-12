package main

import (
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
	caoJSON := `{"nome": "Rex", "raca": "dalmata", "idade": 3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(caoJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)
}
