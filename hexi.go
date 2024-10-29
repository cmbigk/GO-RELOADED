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

func hexa2decimal(words string) int {
	totalValue := 0
	powerofbase := 1

	for i := len(words) - 1; i >= 0; i-- {

		totalValue = totalValue + hexDec(rune(words[i]))*powerofbase
		powerofbase = powerofbase * 16

	}
	return totalValue
}
