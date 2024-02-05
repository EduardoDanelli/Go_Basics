package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

// não precisa de retorno pois a alteração vai ser diretamente na variável
func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 6
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
