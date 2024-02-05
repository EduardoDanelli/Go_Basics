package main

import "fmt"

func main() {

	// sem parâmetro
	func() {
		fmt.Println("Olá Mundo")
	}()

	// com parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("Olá Mundo")

	// com retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Retorno da função")

	fmt.Println(retorno)
}
