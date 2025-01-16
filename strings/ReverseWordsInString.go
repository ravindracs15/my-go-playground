package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	out := ""
	for _, word := range words {
		fmt.Println("@"+word+"@")
		if len(word) == 0 {
			continue
		}
		out = word + " " + out
	}

	return strings.TrimRight(out, " ")
}

func main() {
	fmt.Println(reverseWords("    the sky is blue") + "!!")
	ff := " test"
	fmt.Println(ff[0] == ' ')
}
