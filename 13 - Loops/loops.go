package main

import (
	"fmt"
)

func main() {

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Du", "Dudu", "Edu"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for index, letra := range "Eduardo" {
		fmt.Println(index, string(letra))
	}

	user := map[string]string{
		"nome":      "Jorge",
		"sobrenome": "Silva",
	}

	for index, value := range user {
		fmt.Println(index, value)
	}
}
