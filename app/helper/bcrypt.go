package helper

import "golang.org/x/crypto/bcrypt"

func HassPass(p string) string {
	salt := 10
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}