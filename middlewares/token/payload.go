package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by verifyToken function.
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload is a struct that contains the payload data of the token.
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	IsAdmin   bool      `json:"is_admin"`
	IsUsedAt  time.Time `json:"isused_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload is a function that creates a new token payload with a specific userID and duration.
func NewPayload(userID string, duration time.Duration, isAdmin bool) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IsAdmin:   isAdmin,
		IsUsedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid is a function that checks if the token payload is valid or not.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
