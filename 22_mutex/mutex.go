package main

import (
    "fmt"
    "sync"
)


//1. Problem Without Mutex (Race Condition)
// var counter = 0

// func increment(wg *sync.WaitGroup) {
//     for i := 0; i < 1000; i++ {
//         counter++ // not thread-safe
//     }
//     wg.Done()
// }

// func main() {
//     var wg sync.WaitGroup
//     wg.Add(2)

//     go increment(&wg)
//     go increment(&wg)

//     wg.Wait()
//     fmt.Println("Final Counter:", counter)
// }


//2. Fix with Mutex

// var counter = 0
// var mu sync.Mutex

// func increment(wg *sync.WaitGroup) {
//     for i := 0; i < 1000; i++ {
//         mu.Lock()   // lock before accessing shared data
//         counter++
//         mu.Unlock() // unlock after update
//     }
//     wg.Done()
// }

// func main() {
//     var wg sync.WaitGroup
//     wg.Add(2)

//     go increment(&wg)
//     go increment(&wg)

//     wg.Wait()
//     fmt.Println("Final Counter:", counter)
// }

//3. Using Mutex in a Struc

// type SafeCounter struct {
//     mu sync.Mutex
//     v  map[string]int
// }

// func (c *SafeCounter) Inc(key string) {
//     c.mu.Lock()
//     c.v[key]++
//     c.mu.Unlock()
// }

// func (c *SafeCounter) Value(key string) int {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     return c.v[key]
// }

// func main() {
//     c := SafeCounter{v: make(map[string]int)}

//     var wg sync.WaitGroup
//     for i := 0; i < 1000; i++ {
//         wg.Add(1)
//         go func() {
//             defer wg.Done()
//             c.Inc("somekey")
//         }()
//     }
//     wg.Wait()
//     fmt.Println("Final Value:", c.Value("somekey"))
// }

//5. RWMutex (Read-Write Mutex)	

type SafeData struct {
    mu sync.RWMutex
    data int
}

func (s *SafeData) Read() int {
    s.mu.RLock()         // multiple readers allowed
    defer s.mu.RUnlock()
    return s.data
}

func (s *SafeData) Write(val int) {
    s.mu.Lock()          // only one writer allowed
    s.data = val
    s.mu.Unlock()
}

func main() {
    var wg sync.WaitGroup
    sd := SafeData{}

    // Writer
    wg.Add(1)
    go func() {
        defer wg.Done()
        sd.Write(42)
    }()

    // Reader
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Read Value:", sd.Read())
    }()

    wg.Wait()
}