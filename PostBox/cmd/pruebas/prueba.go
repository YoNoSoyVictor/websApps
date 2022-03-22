package main

import (
	"fmt"
	"net/mail"
)

func ValidateMail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	mail := "holasoyunaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:maiaaaaaaasdawdqwe"
	err := ValidateMail(mail)
	fmt.Println(err)
}
