package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, slat string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s", password, slat)), 15)
	return string(bytes), err
}

func CheckPasswordHash(password, slat, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(fmt.Sprintf("%s%s", password, slat))) == nil
}
