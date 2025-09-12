package main

import "fmt"

// interfaces -> a collection of method signatures
// interfaces are implemented implicitly
// if a type has methods defined in an interface, it implements that interface

// defining an interface
type paymenter interface {
	pay(amount float32)
}

// struct that uses the interface
type payment struct {
	gateway paymenter
}

// method to make payment
func (p payment) makePayment(amount float32){
	p.gateway.pay(amount)
}

// different types that implement the paymenter interface
type phonepe struct {}
func (pp phonepe) pay(amount float32){
	fmt.Println("payement successful form phonepe",amount)
}

type paytm struct {}
func (py paytm) pay(amount float32){
	fmt.Println("payement successful form paytm",amount)
}

type googlepay struct {}
func (gp googlepay) pay(amount float32){
	fmt.Println("payement successful form googlepay",amount)
}

// main function
func main(){
	// creating instances of payment with different gateways
	p1 := payment{
		gateway: phonepe{},
	}
	p2 := payment{
		gateway: paytm{},
	}
	p3 := payment{
		gateway: googlepay{},
	}
	
	// making payments
	p1.makePayment(1000)
	p2.makePayment(2000)
	p3.makePayment(3000)
}
