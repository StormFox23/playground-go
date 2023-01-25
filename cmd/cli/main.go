package main

import (
	"fmt"
	"os/exec"
	"strings"

	arp "github.com/StormFox23/playground-go/internal/arp"
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

	for ip, _ := range arp.Table() {
		fmt.Printf("%s : %s\n", ip, arp.Search(ip))
	}
}

func arp_linux() {
	data, err := exec.Command("arp", "-an").Output()
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		// strip brackets around IP
		ip := strings.Replace(fields[1], "(", "", -1)
		ip = strings.Replace(ip, ")", "", -1)

		mac := fields[3]
		fmt.Println(ip, mac)
	}
}
