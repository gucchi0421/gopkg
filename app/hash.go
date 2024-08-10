package app

import (
	"golang.org/x/crypto/bcrypt"
)

// new hash str
func HashNew(p string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

// check hash str
func HashDiff(inp, orig string) error {
	return bcrypt.CompareHashAndPassword([]byte(orig), []byte(inp))
}
