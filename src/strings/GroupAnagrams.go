package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mapS := make(map[string][]string, 0)
	for _, w := range strs {
		byteArray := []byte(w)
		sort.Slice(byteArray, func(i, j int) bool {
			return byteArray[i] < byteArray[j]
		})
		ss := string(byteArray)
		//	fmt.Println(w + "  " + ss)
		mapS[ss] = append(mapS[ss], w)
	}
	ans := make([][]string, 0)
	for _, val := range mapS {
		ans = append(ans, val)
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
