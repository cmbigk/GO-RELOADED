package main

import "strconv"

func ExtractNumber(s string) int {
	Number, err := strconv.Atoi(string(s[:len(s)-1]))
	IsErrNo(err)
	return Number
}
