package main

import "fmt"

func generate(numRows int) [][]int {
	output := make([][]int, numRows)
	output[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		prevRow := output[i-1]
		for j := 1; j < i; j++ {
			row[j] = prevRow[j-1] + prevRow[j]
		}
		row[i] = 1
		output[i] = row
	}
	return output

}

func main() {
	fmt.Println(generate(3))
	exSlice := []int{1}
	exSlice = append(exSlice, 15)
	fmt.Println(exSlice)
}
