package helpers

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("HashPassword: %w", err)
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsPasswordOK(s string) bool {
	var eightOrMore, number, upper, special bool

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case !(unicode.IsLetter(c) || c == ' ' || unicode.IsNumber(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)):
			return false
		}
	}
	eightOrMore = len(s) >= 8
	return eightOrMore && number && upper && special
}
