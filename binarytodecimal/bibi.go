package main

import "fmt"

func biDec(c rune) int {

	if c == '0' || c == '1' {
		return int(c - '0')
	}
	return -1
}

func binary2decimal(s string) int {
	totalValue := 0
	powerofbase := 1

	for i := len(s) - 1; i >= 0; i-- {

		totalValue = totalValue + biDec(rune(s[i]))*powerofbase
		powerofbase = powerofbase * 2
	}
	return totalValue
}

func main() {
	// Test cases
	testCases := []struct {
		binary   string
		expected int
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"101", 5},
		{"1101", 13},
		{"10001", 17},
		{"11111111", 255},
		{"1100101", 101},
		{"10000000000", 1024}, // Large binary number
	}

	for _, testCase := range testCases {
		result := binary2decimal(testCase.binary)
		if result == testCase.expected {
			fmt.Printf("PASS: BinaryToDecimal(%s) = %d\n", testCase.binary, result)
		} else {
			fmt.Printf("FAIL: BinaryToDecimal(%s) = %d, expected %d\n", testCase.binary, result, testCase.expected)
		}
	}
}
