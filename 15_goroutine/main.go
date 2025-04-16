// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func loop(n int, wg *sync.WaitGroup) {
// 	for i := 0; i < 1000000; i++ {
// 		fmt.Print(n)
// 	}
// 	wg.Done()
// }

// func main() {

// 	start := time.Now()

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go loop(10, &wg)

// 	go loop(20, &wg)

// 	wg.Wait()

// 	elapsed := time.Since(start)

// 	fmt.Print("\n time taken: ", elapsed, "\n")

// }

// // normal function - time taken: 3.340015715s
// // time taken: 418.2532ms

package main

import "fmt"

func main() {
	// // Create a buffered channel with a capacity of 2
	// messages := make(chan string, 2)

	// // Send two messages to the channel
	// messages <- "buffered"
	// messages <- "channel"

	// // Receive and print the messages
	// fmt.Println(<-messages) // o/p -> buffered
	// fmt.Println(<-messages) // o/p -> channel

	s := make(map[string]int)
	s["karthik"] = 15
	s["rakesh"] = 17
	s["mohit"] = 19
	fmt.Print(s)
}
