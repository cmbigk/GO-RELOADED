package main

import "strconv"

func biDec(c rune) int {

	if c == '0' || c == '1' {
		return int(c - '0')
	}
	return -1
}

func binary2decimal(words string) string {
	totalValue := 0
	powerofbase := 1

	for i := len(words) - 1; i >= 0; i-- {

		totalValue = totalValue + biDec(rune(words[i]))*powerofbase
		powerofbase = powerofbase * 2
	}
	return strconv.Itoa(int(totalValue))
}
