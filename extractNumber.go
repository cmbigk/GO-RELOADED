package main

import "strconv"

func ExtractNumber(s string) int {
	// Remove the last character from the string and attempt to convert the rest to an integer.
	Number, err := strconv.Atoi(string(s[:len(s)-1]))

	// Check for any errors during conversion.
	IsErrNo(err)

	// Return the extracted integer.
	return Number
}
