package main

func toUpper(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = char - 32
		}
		result += string(char)
	}
	return result
}
