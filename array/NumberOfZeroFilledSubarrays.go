package main

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	count := 0
    currentStreak := 0
    // Iterate through the array
    for _, num := range nums {
        if num == 0 {
            // Increment current streak of zeros
            currentStreak++
            // Add the current streak count to total count
            count += currentStreak
        } else {
            // Reset streak when a non-zero is encountered
            currentStreak = 0
        }
    }
    return int64(count)
}

func main() {
	input := []int{0, 0, 0, 2, 0, 0}
	fmt.Println(zeroFilledSubarray(input))
}
