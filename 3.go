package main

import (
	"./util"
	"fmt"
)

func main() {
	num := int64(13195)

	fmt.Println("Primes less than", num)

	primes := util.PrimesLessThan(num)

	for i, j := range primes {
		fmt.Println(i, j)
	}

	// fmt.Println(highval)
}
