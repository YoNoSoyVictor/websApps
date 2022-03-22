package validations

import (
	"errors"
)

var (
	ErrShort            = errors.New("validations: too short")
	ErrLong             = errors.New("validations: too long")
	ErrInvalidCharacter = errors.New("validations: invalid character")

	InvalidCharacters = []string{"-", ",", "&", "=", "_", "'", "-", "+", "<", ">", "/", " "}
)

func SanitizeString(input string, characters []string) bool {
	//iterates through a string checking wether or not it contains certain symbols

	for i := 0; i < len(input); i++ {
		for _, char := range characters {

			if string(input[i]) == char {
				return true
			}
		}
	}

	return false
}
