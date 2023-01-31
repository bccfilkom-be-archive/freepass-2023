package configs

import (
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

// GetDSN returns data source name for database connection with specified environment variables from .env file.
func GetDSN() (dsn string) {
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	return
}

// GetTokenSymmetricKey returns symmetric key for token maker with specified environment variables from .env file.
func GetTokenSymmetricKey() (key string) {
	key = os.Getenv("TOKEN_SYMMETRIC_KEY")
	return
}

// GetTokenDuration returns token duration for token maker with specified environment variables from .env file.
func GetTokenDuration() (duration time.Duration, err error) {
	duration, err = time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	return
}

// GetHTTPServerAddress returns HTTP server address for routers connection with specified environment variables from .env file.
func GetHTTPServerAddress() (address string) {
	address = os.Getenv("HTTP_SERVER_ADDRESS")
	return
}
