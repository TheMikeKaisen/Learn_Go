package main

import "fmt"

func main() {
	// Initialize a map
	myMap := map[string]int{"apple": 10, "banana": 20, "cherry": 30}

	// Add a new key-value pair
	myMap["date"] = 40

	// Retrieve and print a value
	fmt.Println("Value for 'apple':", myMap["apple"])

	// Check if a key exists
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Banana exists with value:", value)
	} else {
		fmt.Println("Banana does not exist")
	}

	// Delete a key
	delete(myMap, "cherry")
	fmt.Println("After deletion:", myMap)

	// Iterate over the map
	fmt.Println("Iterating over map:")
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Map length
	fmt.Println("Map length:", len(myMap))
}
