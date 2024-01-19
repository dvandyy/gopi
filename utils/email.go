package utils

import "net/mail"

func CheckValidEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
