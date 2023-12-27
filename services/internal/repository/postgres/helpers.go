package postgres

import (
	"errors"

	"github.com/lib/pq"
)

func isUniqueViolationError(err error) bool {
	var postgresError *pq.Error
	ready := errors.As(err, &postgresError)
	if !ready {
		return false
	}

	return postgresError.Code == "23505"
}
