package class

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	"github.com/pborman/uuid"
)

func Hostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)
}

func UUID() {
	uuidWithHyphen := uuid.NewRandom()
	//uuid := String.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuidWithHyphen)
}

// ToUpper uppercases all the runes in its argument string.
func ToUpper(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}

func Info() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}
	fmt.Println("You are using Go version 1.8 or higher!")
}

func split(in string) ([]string, string) {
	var ret = strings.Split(in, ",")
	x := ToUpper("ciao")
	return ret, x
}
