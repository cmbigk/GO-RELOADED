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

/*
func main() {
	// Test cases
	test1 := []string{"This", "is", "an", "apple"}
	test2 := []string{"I", "have", "a", "house", "and", "a", "umbrella"}
	test3 := []string{"He", "is", "a", "honest", "man"}
	test4 := []string{"This", "is", "not", "an", "empty", "", "sentence"}

	// Testing addAn function
	result1 := ReplaceAWithAn(test1)
	result2 := ReplaceAWithAn(test2)
	result3 := ReplaceAWithAn(test3)
	result4 := ReplaceAWithAn(test4)

	// Print results
	fmt.Println("Test 1:", result1)
	fmt.Println("Test 2:", result2)
	fmt.Println("Test 3:", result3)
	fmt.Println("Test 4:", result4)
}
*/
