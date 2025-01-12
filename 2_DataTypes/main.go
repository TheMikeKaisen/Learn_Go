package main

import "fmt"

func main() {
    // Basic Types
    var age int = 25                       // Integer
    var salary float64 = 75000.50          // Float
    var isEmployed bool = true             // Boolean
    var name string = "John Doe"           // String
    var complexNum complex128 = 3 + 4i     // Complex number

    // Derived Types
    var numbers [3]int = [3]int{10, 20, 30} // Array
    var grades = []int{85, 90, 92}          // Slice
    var skills = map[string]int{            // Map
        "Go": 8,
        "Python": 9,
    }

    type Person struct { // Struct
        FirstName string
        LastName  string
        Age       int
    }
    var person = Person{"Alice", "Smith", 30}

    // Pointers
    var x int = 42
    var ptr *int = &x // Pointer to x

    // Interface Example
    var emptyInterface interface{} = "I can store anything!"

    // Function Type
    add := func(a, b int) int { // Function stored in a variable
        return a + b
    }

    // Print all data types
    fmt.Printf("Name: %s, Age: %d, Salary: %.2f, Employed: %t\n", name, age, salary, isEmployed)
    fmt.Printf("Complex Number: %v\n", complexNum)
    fmt.Printf("Array: %v, Slice: %v\n", numbers, grades)
    fmt.Printf("Map (Skills): %v\n", skills)
    fmt.Printf("Struct (Person): %+v\n", person)
    fmt.Printf("Pointer Value: %d (Address: %p)\n", *ptr, ptr)
    fmt.Printf("Interface Value: %v\n", emptyInterface)
    fmt.Printf("Function Result (Add): %d\n", add(10, 20))
}
