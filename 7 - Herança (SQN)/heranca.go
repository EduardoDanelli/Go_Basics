package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{
		"João",
		"Das Neves",
		28,
		170,
	}

	fmt.Println(p1)

	e1 := estudante{
		p1,
		"Ciências da Computação",
		"UPF",
	}

	fmt.Println(e1)
}
