package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("fail hash password")
	}

	return hashedPassword, nil
}

func CompareHash(dbPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(requestPassword)); err != nil {
		return errors.New("username and password not match")
	}
	return nil
}
