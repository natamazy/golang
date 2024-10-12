package main

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	numbers := strings.Replace(isbn, "-", "", -1)

	if len(numbers) != 10 {
		return false
	}

	var digit byte
	total := 0

	for i := 0; i <= 9; i++ {
		digit = numbers[i]

		if i == 9 && digit == 'X' {
			total += 10
			continue
		}

		if digit < '0' || digit > '9' {
			return false
		}

		total += (10 - i) * int(digit-'0')
	}

	return total%11 == 0
}
