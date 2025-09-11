package main

import (
	"fmt"
	"maps"
)

// maps -> key value pair
// key -> unique
// value -> duplicate
func main(){


	//create a map

	m1 := make(map[string]string)

	m1["name"] = "john"
	m1["age"] = "20"
	m1["city"] = "new york"

	fmt.Println(m1)
	fmt.Println(len(m1))
	fmt.Println(m1["name"])

	//delete function
	delete(m1, "age")
	fmt.Println(m1)
	fmt.Println(len(m1))

	// check if key exists
	value, ok := m1["age"]
	fmt.Println(value)
	fmt.Println(ok)


	//loop through map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// integrer map
	m2 := make(map[string]int)

	m2["math"] = 90
	m2["science"] = 80
	m2["history"] = 70

	fmt.Println(m2)

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// equality of maps

	 m3 := map[string]int{"a": 1, "b": 2}
	 m4 := map[string]int{"a": 1, "b": 3}	
	 fmt.Println(maps.Equal( m3 , m4)) // invalid operation: m3 == m4 (map can only be compared to nil)
}