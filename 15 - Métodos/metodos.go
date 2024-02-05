package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário: %s no banco de dados\n", u.nome)
}

func (u *usuario) fazAniversario() {
	u.idade++
}

func main() {
	user1 := usuario{"Usuário 1", 29}
	fmt.Println(user1)
	user1.salvar()
	user1.fazAniversario()
	fmt.Println(user1.idade)
}
