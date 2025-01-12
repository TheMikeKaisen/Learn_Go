package main

import "fmt"

func main() {

	// using make method
	s1 := make([]int, 5)
	fmt.Println(s1)
	s1 = append(s1, 60, 70)
	fmt.Println(s1)

	s2:= make([]int, 2, 5) //slice of size 2 and capacity 5
	fmt.Println(s2)
	s2 = append(s2, 5, 6, 7, 8, 9, 10)
	fmt.Println(s2)


	// Declare and initialize a slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Original Slice:", slice)
	fmt.Println("size of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))

	// Access elements
	fmt.Println("First Element:", slice[0])
	fmt.Println("Last Element:", slice[len(slice)-1])

	// Append elements
	slice = append(slice, 60, 70)
	fmt.Println("After Append:", slice)

	// Create a slice from an array
	arr := [5]int{100, 200, 300, 400, 500}
	subSlice := arr[1:4]
	fmt.Println("Sub-Slice from Array:", subSlice)

	// Modify a slice (affects the underlying array)
	subSlice[0] = 250
	fmt.Println("Modified Sub-Slice:", subSlice)
	fmt.Println("Modified Array:", arr)

	// Iterating over a slice
	fmt.Println("Iterating:")
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
