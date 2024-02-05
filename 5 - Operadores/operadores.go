package main

import "fmt"

func main() {
	//Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDivisao)

	// NÃO SE PODE SOMAR... TIPOS DE DADOS DIFERENTES = INT16 + INT32 == NÃO PODE
	var n1 int16 = 32
	var n2 int16 = 64
	soma2 := n1 + n2
	fmt.Println(soma2)

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Lógicos
	fmt.Println("---------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 6
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 2
	fmt.Println(numero)
	numero %= 4
	fmt.Println(numero)

}
