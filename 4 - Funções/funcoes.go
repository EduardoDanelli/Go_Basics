package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 23)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Função F")

	var f2 = func(txt string) string {
		return txt
	}

	resultado := f2("Função f2 com retorno")
	fmt.Println(resultado)

	resultadosCalculosSoma, resultadosCalculosSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosCalculosSoma, resultadosCalculosSubtracao)

	resultadosCalculosSoma2, _ := calculosMatematicos(18, 15)
	fmt.Println(resultadosCalculosSoma2)

	_, resultadosCalculosSubtracao2 := calculosMatematicos(34, 15)
	fmt.Println(resultadosCalculosSubtracao2)
}
