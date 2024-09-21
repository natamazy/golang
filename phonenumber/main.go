package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	for _, val := range phoneNumber {
		if !unicode.IsDigit(val) {
			phoneNumber = strings.ReplaceAll(phoneNumber, string(val), "")
		}
	}

	phoneNumberLen := len(phoneNumber)

	if (phoneNumberLen < 10 || phoneNumberLen > 11) || (phoneNumberLen == 11 && phoneNumber[0] != '1') {
		return "", errors.New("not a valid phone number")
	}

	if phoneNumberLen == 11 {
		phoneNumber = phoneNumber[1:]
	}

	if phoneNumber[0] <= '1' || phoneNumber[3] <= '1' {
		return "", errors.New("not a valid phone number")
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	areaCode, _ := AreaCode(phoneNumber)

	return fmt.Sprintf("(%s) %s-%s", areaCode, phoneNumber[3:6], phoneNumber[6:]), nil
}

func main() {
	phoneNumber, _ := Format("513.99.0253")
	fmt.Println(phoneNumber)
}
