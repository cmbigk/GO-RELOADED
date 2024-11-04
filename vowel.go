package main

// IsVowelOrH checks if the provided character is a vowel (a, e, i, o, u) or the letter 'h' (both uppercase and lowercase).
func IsVowelOrH(char rune) bool {
	vowels := "aeiouAEIOUhH"
	for _, each := range vowels {
		if each == char {
			return true
		}
	}
	return false
}

func ReplaceAWithAn(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		// Check if the current word is "a" or "A" and the next word starts with a vowel or 'h'
		if (words[i] == "a" || words[i] == "A") && len(words[i+1]) > 0 && IsVowelOrH(rune(words[i+1][0])) {
			// Replace "a" with "an" or "A" with "An" based on the case
			if words[i] == "a" {
				words[i] = "an"
			} else if words[i] == "A" {
				words[i] = "An"
			}
		}
	}
	return words
}
