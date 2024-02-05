package main

import "fmt"

// função que recebe parâmetro váriatico "..."
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	somaNumeros := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(somaNumeros)

	escrever("Opa", 1, 2, 3, 4, 5, 6)
}
