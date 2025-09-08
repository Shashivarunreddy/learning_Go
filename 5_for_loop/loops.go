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


	// we can use break and continue in for loop
	for k := 1; k <= 5; k++ {
		if k == 3 {
			continue // skip the rest of the loop when k is 3
		}
		fmt.Println(k)
		if k == 4 {
			break // exit the loop when k is 4
		}

	}


	// range loop
	for a := range 4{
		fmt.Println(a)
	}
}