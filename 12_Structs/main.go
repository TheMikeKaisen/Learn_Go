package main

import "fmt"

// Struct definition
type Car struct {
	Brand string
	Model string
	Year  int
}

// Method with a pointer receiver
func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear
}

func main() {
	car := Car{Brand: "Tesla", Model: "Model S", Year: 2021}

	fmt.Println(car) // Output: {Tesla Model S 2021}

	car.UpdateYear(2023) // Call method to update the year
	fmt.Println(car)     // Output: {Tesla Model S 2023}
}
