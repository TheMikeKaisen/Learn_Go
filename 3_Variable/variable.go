package main

import "fmt"

const c = 56

func main() {
	// Variable declaration
	var name string = "Alice" // Explicit type
	var age = 25              // Type inferred
	height := 5.7             // Short declaration

	const c = 5
	fmt.Println(c)

	// Multiple declarations
	var x, y int = 10, 20
	var isStudent, grade = true, 'A'

	// Constants
	const pi = 3.14159

	// Zero value
	var uninitialized int // Defaults to 0

	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
	fmt.Printf("X: %d, Y: %d, Is Student: %t, Grade: %c\n", x, y, isStudent, grade)
	fmt.Printf("Pi: %.2f, Uninitialized: %d\n", pi, uninitialized)
}
