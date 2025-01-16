package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxVal := math.MinInt
	currSum := 0
	for _, num := range nums {
		currSum += num
		if currSum > maxVal {
			maxVal = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxVal
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum Subarray Sum:", maxSubArray(nums))
}
