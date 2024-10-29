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
	for i := 0; i < len(words); i++ {
		for len(words[i]) > 0 && IsPunctuation(rune(words[i][0])) {
			if i == 0 {
				words[i] = words[i][1:]
			} else {
				words[i-1] = words[i-1] + string(words[i][0])
				words[i] = words[i][1:]
			}
		}
	}
	return removeSlice(words)
}
