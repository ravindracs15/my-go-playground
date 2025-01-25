package main

import "fmt"

func singleNumber(nums []int) (res int) {
	for _, n := range nums {
		res = res^n
	}
	return
}

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
}