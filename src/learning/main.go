/*
--- variable declaration

var numbers [5]int // Declares an array named
'numbers' of 5 integers, initialized to zero values (0s).

--other way
var arrayName = [length]dataType{value1, value2, ...}

CSV files: https://articles.wesionary.team/read-and-write-csv-file-in-go-b445e34968e9

---- define functions

	func functionName(parameter1 type1, parameter2 type2) returnType {
	    // Function body - code to be executed
	    return value // Optional, if returnType is specified
	}

https://go.dev/tour/basics/4

package main

import "fmt"

	func add(x int, y int) int {
		return x + y
	}

	func main() {
		fmt.Println(add(42, 13))
	}
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func array_looper(arr []int, val int) {

	for i := 0; i < len(arr); i++ {
		check := arr[i]
		if check == val {
			fmt.Println("Hit")
		} else {
			fmt.Println("Miss")
		}
	}
}

func main() {
	fmt.Println("Ariel Hello")
	array_name := [...]int{10, 2, 3, 422}

	for x := len(array_name) - 1; x >= 0; x-- {
		fmt.Println(array_name[x], " len")
	}

	var new_array = [...]int{1, 2, 3, 4, 50, 5, 6, 7, 100}

	for new_index := 0; new_index <= len(new_array); new_index++ {
		fmt.Println("index: ", new_index)
	}

	fmt.Println(new_array)

	var while_counter int
	while_counter = 0

	while_check := 50

	for while_counter < len(new_array) {
		fmt.Println(while_counter)

		current := new_array[while_counter]

		if current == while_check {
			fmt.Println("Hit")
		} else {
			fmt.Println("NOOOO")
		}
		while_counter++

	}
	var new_array_2 = []int{1, 2, 3, 4, 50, 5, 6, 7, 100}

	array_looper(new_array_2, 4)

	//	read and return from CSV
	//

	file_name := "/Users/arielwhittingham/Downloads/customers-100_random_csv.csv"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(file)

	defer file.Close()

	for x := 50; x <= 100; x++ {
		fmt.Println("num: ", x)
	}

}
