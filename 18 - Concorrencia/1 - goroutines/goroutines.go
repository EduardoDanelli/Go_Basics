package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Pikachu, choque do trovão!!!")
	escrever("Pikachu, eu escolho você!!!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
