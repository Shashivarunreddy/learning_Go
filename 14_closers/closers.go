package main

// closers -> function which references variables from outside its body

func couter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main(){
	fn := couter()
	println(fn())
	println(fn())
	println(fn())

}