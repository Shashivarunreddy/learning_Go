package main

import "fmt"

// functions -> block of code that performs a specific task
// func functionName(parameters) returnType {}
// parameters -> input to the function
// returnType -> output of the function

//simple function
func add(a int, b int) int {
	return a + b
}

// function with multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// function which retuns a function

func outer() func(a int) int {
	return func(a int ) int{
		return a + 1
	}
}

// different types of return types

func f1() (string, int , bool) {
	return "hello", 10, true
}


func main(){

	sum := add(10, 20)
	fmt.Println(sum)

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fn := outer()
	fmt.Println(fn(5))
	fmt.Println(outer()(10)) // calling the returned function directly

	s, n, _ := f1()
	fmt.Println(s, n)

}