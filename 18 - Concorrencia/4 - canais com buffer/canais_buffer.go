package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Opa"
	canal <- "Gurizada"

	mensagem := <-canal
	fmt.Println(mensagem)

}
