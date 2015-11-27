package main


import (
	"fmt"
)

var (
	upperBound = 1000
	runningTotal = 0
)

func main() {
	for i := 1; i < upperBound; i++ {
		if i % 5 == 0 || i % 3 == 0 {
			// fmt.Println(i)
			runningTotal += i
		}
	}

	fmt.Println(runningTotal)
}

// output
// 233168
