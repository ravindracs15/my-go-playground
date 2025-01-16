package main

import "fmt"

func reverseString(s []byte) {
	len := len(s)
	for i := 0; i < len/2; i++ {
		tmp := s[len-i-1]
		s[len-i-1] = s[i]
		s[i] = tmp
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))

}
