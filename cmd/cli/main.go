package main

import (
	"fmt"
	"runtime"
	"time"

	class "github.com/StormFox23/playground-go/internal/class"
	mqtt "github.com/StormFox23/playground-go/internal/mqtt"
)

func main() {
	fmt.Println("Hello, world.")

	mqtt.Connect()
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

	var mem runtime.MemStats
	stats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	stats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	stats(mem)
	fmt.Println("**END**")
}

func stats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}
