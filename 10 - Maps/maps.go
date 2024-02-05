package main

import "fmt"

func main() {

	user := map[string]string{
		"nome":      "Jorge",
		"sobrenome": "Guimarães",
	}

	fmt.Println(user)
	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Fredeslau",
		},
		"curso": {
			"nome":   "Engenharia de Software",
			"campus": "UPF",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["nome"]["primeiro"])
	fmt.Println(user2["curso"]["nome"])

	user2["signo"] = map[string]string{
		"nome": "Áries",
	}

	fmt.Println(user2)
}
