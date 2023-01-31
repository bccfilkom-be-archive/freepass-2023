package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoMaker is a struct that contains the paseto and symmetric key.
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker is a function that creates a new paseto maker instance.
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %v characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{paseto: paseto.NewV2(), symmetricKey: []byte(symmetricKey)}
	return maker, nil
}

// CreateToken is a method that creates a new paseto token for a specific userID and duration.
func (maker *PasetoMaker) CreateToken(userID string, duration time.Duration, isAdmin bool) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration, isAdmin)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken is a method that checks if the paseto token is valid or not.
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
