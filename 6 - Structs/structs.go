package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var user usuario
	user.nome = "Eduardo"
	user.idade = 28

	fmt.Println(user)

	usuario2 := usuario{
		"Edu",
		29,
		endereco{
			"Rua dos bobos", 0,
		},
	}
	fmt.Println(usuario2)

	// funciona passando apenas um dos campos também
	usuario3 := usuario{nome: "ED", idade: 29}
	fmt.Println(usuario3)
}
