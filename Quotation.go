package main

import (
	"strings"
)

func QuoteHandling(input string) string {
	var result string
	inQuotes := false
	var quotedContent string

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == '\'' {
			if inQuotes {
				inQuotes = false
				result += "'" + strings.TrimSpace(quotedContent) + "'"
				quotedContent = ""
			} else {
				inQuotes = true
			}
			continue
		}

		if inQuotes {
			quotedContent += string(char)

		} else {
			result += string(char)
		}
	}
	return strings.TrimSpace(result)
}
