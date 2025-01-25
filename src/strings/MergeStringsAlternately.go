package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	var result []byte
	i, j := 0, 0
	n1, n2 := len(word1), len(word2)

	// Merge characters alternately
	for i < n1 || j < n2 {
		if i < n1 {
			result = append(result, word1[i])
			i++
		}
		if j < n2 {
			result = append(result, word2[j])
			j++
		}
	}

	return string(result)
}

func main() {
	fmt.Println(mergeAlternately("abc", "def566"))
}
