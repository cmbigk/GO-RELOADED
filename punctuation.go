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

func PunPun(s []string) []string {
	!!!

}
