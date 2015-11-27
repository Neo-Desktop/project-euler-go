package main


import (
	"fmt"
)

var (
	upperBound = 4000000
	runningTotal = 0
)

func main() {
    for i, j := 0, 1; j < upperBound; i, j = i+j,i {
    	// fmt.Println(i)

        if i % 2 == 0 {
        	runningTotal += i
        }
    }

	fmt.Println(runningTotal)
}

// output
// 4613732
