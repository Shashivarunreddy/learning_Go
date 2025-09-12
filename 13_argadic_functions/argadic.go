package main

import "fmt"

// argadic functions -> functions which can take variable number of arguments

// func functionName(parameters ...type) returnType {}
// parameters -> input to the function
// returnType -> output of the function

func add(nums ...int) int{
	sum :=0

	for _, n := range nums{
		sum += n
	}
	return sum
}

func main(){

	result := add(1,2,3,4,5)
	fmt.Println(result)

	list := []int{1,2,3,4,5,6}
	result1 := add(list...)
	fmt.Println(result1)

}