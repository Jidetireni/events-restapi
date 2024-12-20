package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwdByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(passwdByte), nil
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
