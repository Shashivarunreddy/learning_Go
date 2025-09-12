package main

import (
	"fmt"
	"sync"
)

// goroutines -> lightweight threads managed by the Go runtime
// goroutines are functions or methods that run concurrently with other functions or methods
// goroutines are created using the "go" keyword
// goroutines are multiplexed onto multiple OS threads
// goroutines are cheaper than threads

func task(i int , wg *sync.WaitGroup) {
	// defer wg.Done() // decrementing the wait group counter when the goroutine completes
	defer wg.Done()
	fmt.Println("Task", i, "is being processed")

}

func main() {

	// creating wait group to wait for all goroutines to finish

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)  // incrementing the wait group counter
		go task(i, &wg) // creating a goroutine
	}

	// waiting for all goroutines to finish
	wg.Wait()
	
}
