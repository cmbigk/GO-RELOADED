package main

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
