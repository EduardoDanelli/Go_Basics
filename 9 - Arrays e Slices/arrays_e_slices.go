package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(array3)

	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice1))

	slice1 = append(slice1, 60)
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 5
	fmt.Println(slice2)

	//Arrays internos
	slice3 := make([]int, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacity

	slice4 := make([]int, 5, 6)
	slice4 = append(slice4, 7)
	slice4 = append(slice4, 8)
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacity
}
