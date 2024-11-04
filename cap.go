package main

func Capitalize(words string) string {
	result := "" // Initialize the result string

	// Loop through each character in the string
	for i, char := range words {
		if i == 0 {
			// Capitalize the first character
			result += toUpper(string(char))
		} else {
			// Convert all subsequent characters to lowercase
			result += toLower(string(char))
		}
	}

	return result // Return the final capitalized string
}
