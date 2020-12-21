package main

import (
	"runtime"
	"testing"
)

func TestSortBinaryArrayBySwapping(t *testing.T) {
	testCases := []struct {
		description string
		array       []int
		expected    []int
	}{
		{
			description: "start with zero",
			array:       []int{0, 1, 1, 0},
			expected:    []int{0, 0, 1, 1},
		},
		{
			description: "start with one",
			array:       []int{1, 1, 1, 0},
			expected:    []int{0, 1, 1, 1},
		},
		{
			description: "empty array",
			array:       []int{},
			expected:    []int{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.description, func(t *testing.T) {
			actual := sortBinaryArrayBySwapping(testCase.array)
			shouldBeEqual(t, len(testCase.expected), len(actual))

			for i := range testCase.expected {
				shouldBeEqual(t, testCase.expected[i], actual[i])
			}
		})
	}

}

func TestSortBinaryArray(t *testing.T) {
	testCases := []struct {
		description string
		array       []int
		expected    []int
	}{
		{
			description: "start with zero",
			array:       []int{0, 1, 1, 0},
			expected:    []int{0, 0, 1, 1},
		},
		{
			description: "start with one",
			array:       []int{1, 1, 1, 0},
			expected:    []int{0, 1, 1, 1},
		},
		{
			description: "empty array",
			array:       []int{},
			expected:    []int{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.description, func(t *testing.T) {
			actual := sortBinaryArray(testCase.array)
			shouldBeEqual(t, len(testCase.expected), len(actual))

			for i := range testCase.expected {
				shouldBeEqual(t, testCase.expected[i], actual[i])
			}
		})
	}

}

func shouldBeEqual(t *testing.T, expected, actual interface{}) {
	skip := 1

	if expected != actual {
		_, fn, line, _ := runtime.Caller(skip)
		t.Fatalf("\nExp: %v\nGot: %v\nLocation: %s:%d", expected, actual, fn, line)
	}
}
