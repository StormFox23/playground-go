package main

import (
	"fmt"

	class "github.com/StormFox23/playground-go/internal/class"
)

func main() {
	fmt.Println("Hello, world.")
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }
	// e.LeavesRemaining()

	//client.HttpCall()
	class.Hostname()
	class.UUID()
	class.Info()
}
