package main

func removeSlice(words []string) []string {
	result := []string{}

	// Iterate through each string in the input slice.
	for _, char := range words {
		if char != "" { // Check if the current string is not empty.
			result = append(result, char) // Add non-empty strings to the result slice.
		}
	}
	return result
}
