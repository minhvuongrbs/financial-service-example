package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pw string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(bs), nil
}

func ComparePassword(hpw, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpw), []byte(pw))
}
