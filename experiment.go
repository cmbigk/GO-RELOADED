package main

import (
	"strings"
)

func experiment(sentence string) string {
	words := strings.Fields(sentence)
	if len(words) > 0 {
		words[0] = Capitalize(words[0])
	}
	var num int
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 {
				words[i-1] = hexa2decimal(words[i-1])
				words[i] = ""
			}
		case "(bin)":
			if i > 0 {
				words[i-1] = binary2decimal(words[i-1])
				words[i] = ""
			}
		case "(up)":
			if i > 0 {
				words[i-1] = toUpper(words[i-1])
				words[i] = ""
			}
		case "(low)":
			if i > 0 {
				words[i-1] = toLower(words[i-1])
				words[i] = ""
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = Capitalize(words[i-1])
				words[i] = ""
			}
		case "(low,":
			if i < len(words)-1 {
				num = ExtractNumber(words[i+1])
				NumModify(num, i, words, toLower)
				words[i] = ""
				words[i+1] = ""
			}
		case "(up,":
			if i < len(words)-1 {
				num = ExtractNumber(words[i+1])
				NumModify(num, i, words, toUpper)
				words[i] = ""
				words[i+1] = ""
			}
		case "(cap,":
			if i < len(words)-1 {
				num = ExtractNumber(words[i+1])
				NumModify(num, i, words, Capitalize)
				words[i] = ""
				words[i+1] = ""
			}
		}
	}
	edited := ReplaceAWithAn(QuoteHandle(Punctuation(removeSlice(words))))
	return strings.Join(edited, " ")
}
