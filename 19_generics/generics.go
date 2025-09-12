package main

import "fmt"

// generics -> a way to write reusable code
// generics allow us to write functions and data structures that can work with any data type
// generics are defined using type parameters
// we can use comparable key or mentions different constraints
// generics can be used with functions, structs, interfaces and methods
func print[T any](s []T){
	for _, v := range s{
		fmt.Println(v)
	}
}
func main(){

	s1 := []int{1,2,3,4,5}
	print(s1)
	s2 := []string{"afheof","befwef","chrh","dwrqw","dwffge"}
	print(s2)
	s3 := []float64{1.1,2.2,3.3,4.4,5.5}
	print(s3)
}