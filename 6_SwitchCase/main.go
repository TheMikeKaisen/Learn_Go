package main

import "fmt"

func main() {
    day := 3

    // Basic switch statement
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    default:
        fmt.Println("Invalid day")
    }

    // Switch without expression (acts like if-else)
    num := 15
    switch {
    case num%2 == 0:
        fmt.Println("Even number")
    case num%3 == 0:
        fmt.Println("Divisible by 3")
    default:
        fmt.Println("Odd and not divisible by 3")
    }

    // Switch with fallthrough
    grade := "B"
    switch grade {
    case "A":
        fmt.Println("Excellent")
        fallthrough
    case "B":
        fmt.Println("Good")
        fallthrough
    case "C":
        fmt.Println("Average")
    default:
        fmt.Println("Fail")
    }
}
