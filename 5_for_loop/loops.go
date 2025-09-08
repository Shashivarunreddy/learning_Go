package main

import "fmt"

// there is only one type of loop in go: for loop
// for loop can be used as while loop and foreach loop
// there is no while loop and foreach loop in go
func main(){

    // while loop using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	
	// traditional for loop
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	//infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

}