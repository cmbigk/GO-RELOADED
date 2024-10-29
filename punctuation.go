package main

import "fmt"

func IsPunctuation(char rune) bool {

	Punlay := ".,!?;:"
	for _, each := range Punlay {
		if each == char {
			return true
		}
	}
	return false
}

func PunPun(words []string) []string {
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
	return words
}

func main() {
	// Test Case 1: Simple punctuation removal at the start of the first word
	test1 := []string{"...Hello", "world!"}
	fmt.Println("Test 1:", PunPun(test1))
	// Expected: ["Hello", "world!"]

	// Test Case 2: Punctuation at the beginning of multiple words
	test2 := []string{"Hello", ",world", "...Nice", "to", "see", "you!"}
	fmt.Println("Test 2:", PunPun(test2))
	// Expected: ["Hello,", "world!", "Nice", "to", "see", "you!"]

	// Test Case 3: Consecutive words with punctuation in the middle of a sentence
	test3 := []string{"Hello", ",", "world!"}
	fmt.Println("Test 3:", PunPun(test3))
	// Expected: ["Hello,", "world!"]

	// Test Case 4: Words with mixed punctuation at the beginning
	test4 := []string{"...Hello", ",world", "!!Nice", "to", "see", "...you!"}
	fmt.Println("Test 4:", PunPun(test4))
	// Expected: ["Hello,", "world", "Nice", "to", "see", "you!"]

	// Test Case 5: No punctuation at the beginning of any word
	test5 := []string{"Hello", "world", "this", "is", "a", "test"}
	fmt.Println("Test 5:", PunPun(test5))
	// Expected: ["Hello", "world", "this", "is", "a", "test"]

	// Test Case 6: Only punctuation words
	test6 := []string{"...", "!!", ",,,", "!?"}
	fmt.Println("Test 6:", PunPun(test6))
	// Expected: []

	// Test Case 7: Mixed words and punctuation words
	test7 := []string{"Hello", "...", "world", "!!", "Nice", "to", "see", "you!"}
	fmt.Println("Test 7:", PunPun(test7))
	// Expected: ["Hello", "world", "Nice", "to", "see", "you!"]

	// New Test Case 8: Sentence with misplaced punctuation and extra spaces
	test8 := []string{"I", "was", "sitting", "over", "there", ",and", "then", "BAMM", "!!"}
	fmt.Println("Test 8:", PunPun(test8))
	// Expected: ["I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"]

	// New Test Case 9: Sentence with ellipsis and spaces
	test9 := []string{"I", "was", "thinking", "...", "You", "were", "right"}
	fmt.Println("Test 9:", PunPun(test9))
	// Expected: ["I", "was", "thinking...", "You", "were", "right"]
}
