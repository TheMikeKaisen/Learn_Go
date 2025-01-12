package main

import "fmt"

type OrderStatus int

const (
	Received  OrderStatus = iota // 0
	Delivered                    // 1
	Pending                      // 2
	Process                      // 3
)

func CheckOrderStatus(status OrderStatus) {
	fmt.Println("status: ", status)
}
func main() {
	CheckOrderStatus(Process)
}
