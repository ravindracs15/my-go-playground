package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	mapEmails := map[string]bool{}
	for _, email := range emails {
		parts := strings.Split(email, "@")
		part0Array := strings.Split(parts[0], "+")
		mapEmails[strings.ReplaceAll(part0Array[0], ".", "")+"@"+parts[1]] = true
	}
	return len(mapEmails)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
