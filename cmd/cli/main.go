package main

import (
	"fmt"

	employee "github.com/StormFox23/playground-go/internal/employee"
	client "github.com/StormFox23/playground-go/internal/http"
)

func main() {
	fmt.Println("Hello, world.")
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

	client.HttpCall()

}
