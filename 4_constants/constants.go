package main

import "fmt"

const name2 string = "Gophersss"
const pi = 3.14

// we cannot use := for constants
//we cannot change the value of a constant
func main() {
	const name string = "Gopher"
	const age = 10


	const(
		a = 1
		name3 = "Gopher3"
		age3 = 20
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name2)
	fmt.Println(pi)
	fmt.Println(a)
	fmt.Println(name3)
	fmt.Println(age3)
}