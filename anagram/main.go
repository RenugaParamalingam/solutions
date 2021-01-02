package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagram("ac", "bb"))
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sortedS := make([]string, 0)
	sortedT := make([]string, 0)

	for _, r := range s {
		sortedS = append(sortedS, string(r))
	}

	sort.Strings(sortedS)

	for _, r := range t {
		sortedT = append(sortedT, string(r))
	}

	sort.Strings(sortedT)

	for i := range sortedS {
		if sortedS[i] != sortedT[i] {
			return false
		}
	}

	return true
}
