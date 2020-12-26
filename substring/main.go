package main

import (
	"fmt"
)

func main() {

	fmt.Println("mySubString: ", mySubString("abca"))
	fmt.Println("subString: ", subString("abca"))
	fmt.Println("longestSubstringWithoutRepeatingChar: ", longestSubstringWithoutRepeatingChar("pwwkew"))
}

// my logic
func mySubString(s string) []string {
	subString := make([]string, 0)

	l := len(s)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			subString = append(subString, s[i:j+1])
		}

	}
	return subString
}

// another approach got from internet
func subString(s string) []string {
	subStr := make([]string, 0)

	getSubstring := func(s string, i, j int) []string {
		sub := make([]string, 0)
		for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
			sub = append(sub, s[i:j+1])
		}
		return sub
	}

	for i := range s {
		subStr = append(subStr, getSubstring(s, i, i)...)
		subStr = append(subStr, getSubstring(s, i, i+1)...)
	}

	return subStr
}

func longestSubstringWithoutRepeatingChar(s string) int {
	m := make(map[rune]int)
	dupCharIndex, longestLength := 0, 0

	for i, val := range s {
		if mIndex, ok := m[val]; ok {
			dupCharIndex = max(dupCharIndex, mIndex+1)
		}

		m[val] = i
		longestLength = max(longestLength, i-dupCharIndex+1)
	}
	return longestLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
