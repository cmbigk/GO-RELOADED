package main

func Capitalize(words string) string {
	result := ""

	for i, char := range words {
		if i == 0 {
			result += toUpper(string(char))
		} else {
			result += toLower(string(char))
		}
	}

	return result
}
