package main

import "strconv"

func biDec(c rune) int {
	// Check if character is '0' or '1'
	if c == '0' || c == '1' {
		return int(c - '0') // Convert character to integer (0 or 1)
	}
	return -1 // for invalid characters
}

// binary2decimal converts a binary string to its decimal (base 10) representation.
func binary2decimal(words string) string {
	totalValue := 0
	powerofbase := 1

	// Loop through the binary string from the last character to the first
	for i := len(words) - 1; i >= 0; i-- {
		// Convert current character to decimal and multiply by its power of 2
		totalValue = totalValue + biDec(rune(words[i]))*powerofbase
		// Update power of base for the next position (2^1, 2^2, etc.)
		powerofbase = powerofbase * 2
	}

	return strconv.Itoa(totalValue) // Convert the final result to string and return it
}
