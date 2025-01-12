package main

import "fmt"

func main() {
    // Basic for loop
    for i := 1; i <= 3; i++ {
        fmt.Println("Basic Loop:", i)
    }

    // While-like loop
    n := 1
    for n < 3 {
        fmt.Println("While-like Loop:", n)
        n++
    }

    // Infinite loop
    count := 0
    for {
        fmt.Println("Infinite Loop:", count)
        count++
        if count == 3 {
            break
        }
    }

    // Range loop (slice)
    nums := []int{10, 20, 30}
    for idx, val := range nums {
        fmt.Printf("Range Loop - Index: %d, Value: %d\n", idx, val)
    }
}
