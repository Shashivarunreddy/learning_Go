package main

import "fmt"

// enums -> a collection of named constants
// enums are used to define a set of related constants
// enums are defined using the "iota" keyword for auto incrementing values

// defining an enum for the week (int type)
type week int
const(
	week1 week = iota + 1
	week2
	week3
	week4
)

func getWeek(w week){
	fmt.Println("week is", w)
}

// defining an enum for days of the week (string type)
type day string

const (
	sunday    day = "sunday"
	monday        = "monday"
	tuesday       = "tuesday"
	wednesday     = "wednesday"
	thursday      = "thursday"
	friday        = "friday"
	saturday      = "saturday"
)


func getDay(d day){
	fmt.Println("day is", d)
}

func main() {

	getDay(saturday)
	getWeek(week3)

}
