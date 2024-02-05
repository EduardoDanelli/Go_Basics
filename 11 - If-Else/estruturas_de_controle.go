package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// IF INIT
	if numero2 := numero; numero2 > 0 {
		fmt.Println("numéro é maior que zero")
	}
}
