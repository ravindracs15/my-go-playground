package main

import "fmt"

// Function that accepts an array of size 3
func processArray(arr []int) {
    fmt.Println("Array:", arr)
}

func main() {
    arr := []int{1, 2, 3}
    processArray(arr) // Works fine

    // Uncommenting the following line will cause a compilation error
     arr2 := []int{1, 2, 3, 4}
     processArray(arr2) // ERROR: mismatched types
}
