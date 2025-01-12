package main

import "fmt"

func main() {
	// Declare and initialize an array
	arr := [5]int{10, 20, 30, 40, 50}

	// Access elements
	fmt.Println("First Element:", arr[0])
	fmt.Println("Last Element:", arr[len(arr)-1])

	// Modify an element
	arr[2] = 35
	fmt.Println("Modified Array:", arr)

	// Iterate using for loop
	fmt.Println("Using for loop:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	// Iterate using range
	fmt.Println("Using range:")
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Copying arrays
	copyArr := arr // Creates a copy
	copyArr[0] = 100
	fmt.Println("Original Array:", arr)
	fmt.Println("Copied Array:", copyArr)
}
