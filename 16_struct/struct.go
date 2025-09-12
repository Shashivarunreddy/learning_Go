package main

import (
	"fmt"
)

// struct -> collection of fields
// used to group data together to form a record
// struct is a user defined type

type student struct{
	name  string
	age   int
	grade int
}

// struct imbedding -> embedding one struct into another struct

type kids struct {
	name string
	school string
}

type parent struct {
	name string
	age  int
	occupation string
	kids
}

// we can also define methods on structs

type order struct {
	id     int
	amount float64
	status string
}

func newOrder(id int, amount float64, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,

}
    return &myOrder
}

// method to update order status
func (o *order) updateStatus(status string){
	o.status = status
}
// method to display order details
func (o order) display() int {
	return o.id
}


func main(){

	s1 := student{name: "John", age: 20, grade: 90}

	fmt.Println(s1)
	fmt.Println("Name:", s1.name)
	fmt.Println("Age:", s1.age)
	fmt.Println("Grade:", s1.grade)

	s1.age = 30
	fmt.Println("Updated Age:", s1.age)

	// direct initialization and declaration

	language := struct{
		name string
		year int
	}{"Go",2009,}
	fmt.Println(language)


	// struct imbedding
	p1 := parent{
		name: "Alice",
		age:  40,
		occupation: "Engineer",
		kids: kids{
			name:   "Bob",
			school: "ABC School",
		},
	}

	fmt.Println(p1)
	fmt.Println("Parent Name:", p1.name)
	fmt.Println("Kid Name:", p1.kids.name)
	fmt.Println("Kid School:", p1.kids.school)


	o1 := newOrder(1, 100.50, "Processing")
	fmt.Println("Order ID:", o1.display())

	fmt.Println("Order Status:", o1.status)
	fmt.Println(o1)
	o1.updateStatus("Shipped")
	fmt.Println("Updated Order Status:", o1.status)
	fmt.Println(o1)
}