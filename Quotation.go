package main

import "strings"

func QuoteHandle(words []string) []string {
	var QuoteCount int
	for i := 0; i < len(words); i++ {

		// Check for an opening single quote
		// - Only applies if we are currently outside of a quoted phrase (QuoteCount % 2 == 0)
		// - Ensures there's another word following the current one (i < len(words) - 1)
		// - Checks if the current word contains a single quote
		if QuoteCount%2 == 0 && i < len(words)-1 && strings.Contains("'", words[i]) {
			words[i+1] = "'" + words[i+1]
			words[i] = ""
			QuoteCount++

			// Check for a closing single quote
			// - Only applies if we are within a quoted phrase (QuoteCount % 2 == 1)
			// - Ensures there's a previous word to attach the single quote to (i > 0)
			// - Checks if the current word contains a single quote
		} else if QuoteCount%2 == 1 && i > 0 && strings.Contains("'", words[i]) {
			words[i-1] = words[i-1] + "'"
			words[i] = ""
			QuoteCount++
		}
	}
	return removeSlice(words)
}
