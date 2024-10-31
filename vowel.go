package main

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
		if words[i] == "a" && len(words[i+1]) > 0 && IsVowelOrH(rune(words[i+1][0])) {

			words[i] = "an"
		}
	}
	return words
}
