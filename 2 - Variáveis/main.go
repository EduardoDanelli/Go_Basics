package main

import "fmt"

func main() {
	var variavel_string string = "string"
	fmt.Println(variavel_string)

	variavel_2_string := "string 2" // inferência de tipo
	fmt.Println(variavel_2_string)

	var (
		variavel_3_string string = "variavel 3"
		variavel_4_string string = "variavel 4"
	)

	fmt.Println(variavel_3_string, variavel_4_string)

	var5, var6 := "var5", "var6"

	fmt.Println(var5, var6)

	const constante1 string = "const1"
	fmt.Println(constante1)

	// inversão de valores de variáveis
	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}
