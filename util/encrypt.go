package util

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	if encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(encryptedPassword), nil
	}
}

func VerifyPassword(password string, encryptedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}
