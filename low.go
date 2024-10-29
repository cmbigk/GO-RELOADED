package main

func toLower(words string) string {
	result := ""
	for _, char := range words {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		result += string(char)
	}
	return result
}
