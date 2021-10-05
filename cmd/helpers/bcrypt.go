package helpers

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func ComparePassword(pass string, hashedPass string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {
		return false
	}
	return true
}
