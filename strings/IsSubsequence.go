package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) < len(s) {
		return false
	}
	srcIndex := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[srcIndex] {
			srcIndex++
			if srcIndex == len(s) {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
