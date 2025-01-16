package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	firstMax := math.MaxInt32
	secMax := math.MaxInt32
	for _, num := range nums {
		if num <= firstMax {
			firstMax = num
		} else if num <= secMax {
			secMax = num
		} else {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
