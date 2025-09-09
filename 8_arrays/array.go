package main

import "fmt"

func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false

	//int
	var num [5]int
	fmt.Println(num[0])
	num[0] = 10
	fmt.Println(num[0])
	fmt.Println(len(num))
	fmt.Println(num)

	//string
	var names [6]string
	names[0] = "abc"
	names[1] = "def"
	names[2] = "ghi"
	fmt.Println(names)
	fmt.Println(len(names))

	// bool
	var flags [3]bool
	flags[2] = true
	fmt.Println(flags)
	fmt.Println(len(flags))


	// array with declaration and initialization
	var nums = [3]int{1, 2, 3}
	fmt.Println(nums)
	fmt.Println(len(nums))

	// 2D array
	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
	fmt.Println(len(matrix))
	fmt.Println(len(matrix[0])) // number of columns	
}