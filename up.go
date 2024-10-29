package main

func toUpper(words string) string {
	result := ""
	for _, char := range words {
		if char >= 'a' && char <= 'z' {
			char = char - 32
		}
		result += string(char)
	}
	return result
}
