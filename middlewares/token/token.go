package token

import (
	"log"
	"time"

	"github.com/bagashiz/freepass-2023/configs"
)

var (
	tokenMaker    Maker
	tokenDuration time.Duration
	err           error
)

// SetupTokenAuthMiddleware sets up the token authentication middleware.
func SetupTokenAuthMiddleware() {
	tokenMaker, err = NewPasetoMaker(configs.GetTokenSymmetricKey())
	if err != nil {
		log.Fatal(err)
	}

	tokenDuration, err = configs.GetTokenDuration()
	if err != nil {
		log.Fatal(err)
	}
}

// GetTokenMaker returns the token maker.
func GetTokenMaker() Maker {
	return tokenMaker
}

// GetTokenDuration returns the token duration.
func GetTokenDuration() time.Duration {
	return tokenDuration
}
