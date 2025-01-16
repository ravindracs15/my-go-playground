package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if length == 0 {
				continue
			} else {
				break
			}
		}
		length++
	}

	return length
}

func main() {
	fmt.Println(lengthOfLastWord("rr rr15    "))
}
