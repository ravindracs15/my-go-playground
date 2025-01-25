package main

import "fmt"

func areMapsEqual(map1, map2 map[byte]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exists := map2[key]; !exists || value1 != value2 {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}
	mapAngaram := map[byte]int{}
	mapInput := map[byte]int{}
	targetLen := len(p)
	for i := 0; i < targetLen; i++ {
		mapAngaram[p[i]] = mapAngaram[p[i]] + 1
		mapInput[s[i]] = mapInput[s[i]] + 1
	}
	for i := 0; i < len(s); i++ {
		if areMapsEqual(mapAngaram, mapInput) {
			result = append(result, i)
		}
		mapInput[s[i]] = mapInput[s[i]] - 1
		if mapInput[s[i]] == 0 {
			delete(mapInput, s[i])
		}
		if i+targetLen < len(s) {
			mapInput[s[i+targetLen]] = mapInput[s[i+targetLen]] + 1
		}

	}
	return result
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
