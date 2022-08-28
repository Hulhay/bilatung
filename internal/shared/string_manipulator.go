package shared

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func CheckPassword(password, passwordDB string) error {

	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func RemovePunctuation(text string) (string, error) {

	reg, err := regexp.Compile("[a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	filteredText := reg.ReplaceAllString(text, "")

	return filteredText, nil
}
