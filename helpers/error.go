package helpers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// IsUniqueViolation is a function to check if the error is a unique violation error.
func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

// IsForeignKeyViolation is a function to check if the error is a foreign key violation error.
func IsForeignKeyViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23503"
}

// IsValueTooLong is a function to check if the error is value too long error.
func IsValueTooLong(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "22001"
}

// IsInvalidData is a function to check if the error is an invalid data error.
func IsInvalidData(err error) bool {
	return errors.Is(err, gorm.ErrInvalidData)
}

// IsRecordNotFound is a function to check if the error is a record not found error.
func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
