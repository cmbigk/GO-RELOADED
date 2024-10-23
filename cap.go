package main

func Capitalize(s string) string {
	result := ""

	for i, char := range s {
		if i == 0 {
			result += toUpper(string(char))
		} else {
			result += toLower(string(char))
		}
	}

	return result
}
