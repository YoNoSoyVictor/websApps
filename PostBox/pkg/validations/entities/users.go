package validations

import (
	"net/mail"

	"github.com/YoNoSoyVictor/ThoughtBox/pkg/validations"
)

func ValidateMail(email string) error {
	//checks wether an email is valid or not

	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil
}

func ValidateName(name string) error {
	//checks wether a name is valid or not

	if len(name) < 3 {
		return validations.ErrShort
	}

	if len(name) > 20 {
		return validations.ErrLong
	}

	res := validations.SanitizeString(name, validations.InvalidCharacters)
	if res {
		return validations.ErrInvalidCharacter
	}

	return nil
}

func ValidatePassword(password string) error {
	//checks password lenght
	if len(password) < 5 {
		return validations.ErrShort
	}
	return nil
}
