package main

func removeSlice(word []string) []string {
	result := []string{}
	for _, char := range word {
		if char != "" {
			result = append(result, char)
		}
	}
	return result
}
