package main

func Up(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		result += string(char)
	}
	return result
}
