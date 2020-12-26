package main

import (
	"fmt"
)

func main() {

	fmt.Println(mySubString("abca"))
	fmt.Println(subString("abca"))
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
