package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Opa Mundinho bão")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("eduardodanelli@gmail.com")
	fmt.Println(erro)
}
