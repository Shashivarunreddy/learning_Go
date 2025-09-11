package main

import "fmt"
// range -> loop through array, slice, map, string
// for i := range collection
// for i, v := range collection
// for _, v := range collection
func main(){

	nums := []int{10, 20, 30, 40, 50}

	for i := range nums {
		fmt.Println(i)
	}

	for i, v := range nums {
		fmt.Println(i, v)
	}

	for _, v := range nums {
		fmt.Println(v)
	}

	str := "hello"
	for i, v := range str {
		fmt.Println(i, string(v))
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	
}