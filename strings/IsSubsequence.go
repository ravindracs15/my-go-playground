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
func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for j := 1; j < len(strs); j++ {
		s := strs[j]
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {

		}
		p = p[:i]
	}
	return p
}

func main() {
	s := "abc"
	t := "abgdc"
	fmt.Println(isSubsequence(s, t))
	fmt.Println(longestCommonPrefix([]string{s, t}))
}
