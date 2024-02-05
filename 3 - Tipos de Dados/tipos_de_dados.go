package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint = 100000
	fmt.Println(numero2)

	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234545454.45
	fmt.Println(numeroReal3)

	var str string = "string"
	fmt.Println(str)

	str2 := "Uhum"
	fmt.Println(str2)

	// aspas simples pega o número na tabela ASCCI
	char := 'E'
	fmt.Println(char)

	// variáveis possuem valor 0, onde inicia-se
	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
