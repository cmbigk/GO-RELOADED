package main

func IsPunctuation(char rune) bool {
	Punlay := ".,!?;:"

	for _, each := range Punlay {
		if each == char {
			return true
		}
	}
	return false
}

func Punctuation(words []string) []string {
	// Loop through each word in the slice.
	for i := 0; i < len(words); i++ {
		// Continue removing punctuation from the start of the word while it exists.
		for len(words[i]) > 0 && IsPunctuation(rune(words[i][0])) {
			if i == 0 {
				// If it's the first word, remove the leading punctuation directly.
				words[i] = words[i][1:]
			} else {
				// For subsequent words, append the punctuation to the previous word and remove it from the current word.
				words[i-1] = words[i-1] + string(words[i][0])
				words[i] = words[i][1:]
			}
		}
	}
	return removeSlice(words)
}
