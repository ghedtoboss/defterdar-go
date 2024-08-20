package helpers

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(passHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
