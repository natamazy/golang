package main

import "fmt"

func RotationalCipher(plain string, shiftKey int) string {
	ciphered := make([]rune, len(plain))

	for i, val := range plain {
		switch {
		case val >= 65 && val <= 90:
			ciphered[i] = ((val-65+rune(shiftKey))%26 + 65)
		case val >= 97 && val <= 122:
			ciphered[i] = ((val-97+rune(shiftKey))%26 + 97)
		default:
			ciphered[i] = rune(plain[i])
		}
	}

	return string(ciphered)
}

func main() {
	fmt.Println(RotationalCipher("ab_cz", 1))
}
