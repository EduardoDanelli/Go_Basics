package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generica("string")
	generica(3)
	generica(true)
}
