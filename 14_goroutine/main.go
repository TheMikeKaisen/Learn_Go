package main

import (
	"fmt"
	"time"
)

func loop(n int) {
	for i := 0; i < 1000000; i++ {
		fmt.Print(n)
	}
}

func main() {

	start := time.Now()

	loop(10)

	loop(20)

	elapsed := time.Since(start)

	fmt.Print("\n time taken: ", elapsed, "\n")

}

// 