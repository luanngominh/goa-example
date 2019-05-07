package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

//BcryptHashPassword hash password using bcrypt with password ans salt
func BcryptHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

//BcryptCheckPassword check password and passwordDigest is match each other
func BcryptCheckPassword(password, passwordDigest string) bool {
	byteHash := []byte(passwordDigest)
	
	if err := bcrypt.CompareHashAndPassword([]byte(byteHash), []byte(password)); err != nil {
		return false
	}
	
	return true
}