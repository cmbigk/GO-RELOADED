package main

func removeSlice(words []string) []string {
	result := []string{}
	for _, char := range words {
		if char != "" {
			result = append(result, char)
		}
	}
	return result
}
