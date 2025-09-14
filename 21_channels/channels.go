package main

import (
	"fmt"
	"sync"
	// "time"
)

// channels -> a way to communicate between goroutines
// channels are typed conduits through which you can send and receive values with the channel operator, <-
// channels are created using the make function
// channels can be buffered or unbuffered
// unbuffered channels block the sending goroutine until another goroutine receives from the channel
// buffered channels block the sending goroutine only when the buffer is full
// channels are safe for concurrent use by multiple goroutines

// 1. Creating and Using a Channel
// func main(){
// 	// creating a channel
// 	ch := make(chan int)

// 	// sending a value to the channel
// 	go func() {
// 		ch <- 42
// 	}()

// 	// receiving a value from the channel
// 	val := <- ch
// 	fmt.Println(val) // printing the value
// }

//2. Buffered Channels

// func main() {
//     ch := make(chan string, 2) // buffered channel of size 2

//     ch <- "Hello"
//     ch <- "World"
//     // ch <- "Extra" // ❌ would block (buffer full)

//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
// }

//3. Channel Synchronization

// func worker(done chan bool) {
//     fmt.Println("Working...")
//     time.Sleep(2 * time.Second)
//     fmt.Println("Done")
//     done <- true
// }

// func main() {
//     done := make(chan bool)
//     go worker(done)

//     <-done // wait until worker sends signal
// }

//4. Directional Channels

// func send(ch chan<- int) {
//     ch <- 10 // only sending
// }

// func receive(ch <-chan int) {
//     fmt.Println("Received:", <-ch) // only receiving
// }

// func main() {
//     ch := make(chan int)

//     go send(ch)
//     receive(ch)
// }

//5. Closing a Channel

// func main() {
//     ch := make(chan int, 2)
//     ch <- 1
//     ch <- 2
//     close(ch)

//     for val := range ch {
//         fmt.Println(val)
//     }
// }

// 6. Select with Channels

// func main() {
//     ch1 := make(chan string)
//     ch2 := make(chan string)

//     go func() {
//         time.Sleep(1 * time.Second)
//         ch1 <- "Message from ch1"
//     }()
//     go func() {
//         time.Sleep(2 * time.Second)
//         ch2 <- "Message from ch2"
//     }()

//     for i := 0; i < 2; i++ {
//         select {
//         case msg1 := <-ch1:
//             fmt.Println(msg1)
//         case msg2 := <-ch2:
//             fmt.Println(msg2)
//         }
//     }
// }

//7. Default Case in Select

// func main() {
//     ch := make(chan int)
// 	// go func(){
// 	// 	ch <- 5

// 	// }()
//     select {
//     case val := <-ch:
//         fmt.Println("Received:", val)
//     default:
//         fmt.Println("No value received (non-blocking)")
//     }
// }

// 8. Fan-out, Fan-in (Multiple Goroutines with Channels)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- job * 2
        fmt.Printf("Worker %d processed job %d\n", id, job)
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    var wg sync.WaitGroup

    // start 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    close(results)

    for result := range results {
        fmt.Println("Result:", result)
    }
}