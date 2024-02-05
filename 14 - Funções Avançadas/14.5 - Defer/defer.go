package main

import "fmt"

func funcao1() {
	fmt.Println("Exec func1")
}

func funcao2() {
	fmt.Println("Exec func2")
}

func alunoAprovado(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	} else {
		return false
	}
}

func main() {
	defer funcao1() // --- Defer = Adiar
	funcao2()

	fmt.Println(alunoAprovado(7, 9))
}
