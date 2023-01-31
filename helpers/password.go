package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword is a function that returns the bcrypt hash of a password.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return string(hashedPassword), nil
}

// ComparePassword is a function that compares the given password with the hashed password with bcrypt.
func ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
