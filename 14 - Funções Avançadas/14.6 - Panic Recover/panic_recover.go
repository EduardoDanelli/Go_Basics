package main

import "fmt"

func recuperarExec() {
	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso")
	}
}

func alunoAprovado(n1, n2 float64) bool {

	defer recuperarExec()

	media := (n1 + n2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}
	panic("A média é igual à 7")
}

func main() {
	fmt.Println(alunoAprovado(7, 7))
	fmt.Println("Pós exec")
}
