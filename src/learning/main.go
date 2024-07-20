package main

import "fmt"

func main() {
	fmt.Println("Ariel Hello")
	array_name := [...]int{10, 2, 3, 422}

	for x := len(array_name) - 1; x >= 0; x-- {
		fmt.Println(array_name[x], " len")
	}
	fmt.Println(array_name)

}
