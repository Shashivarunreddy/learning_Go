package main

// pointers -> variable that stores the memory address of another variable

// & -> address of operator
// * -> dereference operator

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)

	// fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)

}