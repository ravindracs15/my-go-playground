package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	currMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = currMax
		if currMax < tmp {
			currMax = tmp
		}
	}
	return arr
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	replaceElements(arr)
	fmt.Println(arr)
}
