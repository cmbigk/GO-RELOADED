package main

import "strconv"

func hexDec(s rune) int {
	// Check if character is uppercase hexadecimal ('A'-'F') and convert to decimal.
	if s >= 'A' && s <= 'F' {
		return int(s - 'A' + 10)
	}

	// Check if character is lowercase hexadecimal ('a'-'f') and convert to decimal.
	if s >= 'a' && s <= 'f' {
		return int(s - 'a' + 10)
	}

	// Check if character is numeric ('0'-'9') and convert to decimal.
	if s >= '0' && s <= '9' {
		return int(s - '0')
	}

	return -1
}

func hexa2decimal(words string) string {
	totalValue := 0
	powerofbase := 1

	// Loop through the string from the last character to the first.
	for i := len(words) - 1; i >= 0; i-- {
		// Add the decimal value of the current hex character, scaled by the current power of 16.
		totalValue = totalValue + hexDec(rune(words[i]))*powerofbase

		// Increase power of 16 for the next character to the left.
		powerofbase = powerofbase * 16
	}

	// Convert the total decimal value to a string and return it.
	return strconv.Itoa(int(totalValue))
}
