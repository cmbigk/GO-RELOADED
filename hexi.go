package main

func hexDec(s rune) int {
	if s >= 'A' && s <= 'F' {
		return int(s - 'A' + 10)
	}

	if s >= 'a' && s <= 'f' {
		return int(s - 'a' + 10)
	}

	if s >= '0' && s <= '9' {
		return int(s - '0')
	}

	return -1
}

func hexa2decimal(s string) int {
	totalValue := 0
	powerofbase := 1

	for i := len(s) - 1; i >= 0; i-- {

		totalValue = totalValue + hexDec(rune(s[i]))*powerofbase
		powerofbase = powerofbase * 16

	}
	return totalValue
}
