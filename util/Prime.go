package util

import (
	"math"
)

func PrimesLessThan(highnum int64) []int64 {

	highnum = int64(math.Sqrt(float64(highnum)))
	highnum = highnum + 1

	nums := make([]bool, highnum)
	totalNum := 1

	for i, _ := range nums {
		nums[i] = true
	}

	nums[0] = false
	nums[1] = false

	for i := int64(2); i <= highnum; i++ {
		for j := int64(2); j <= highnum; j++ {
			if i*j < highnum {
				nums[i*j] = false
			}
		}
	}

	for _, b := range nums {
		if b == true {
			totalNum++
		}
	}

	out := make([]int64, totalNum)
	outcount := int64(0)
	for num, b := range nums {
		if b == true {
			out[outcount] = int64(num)
			outcount++
		}
	}

	return out
}
