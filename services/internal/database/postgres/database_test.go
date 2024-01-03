package postgres

import (
	"testing"

	"microservices/services/internal/config"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectWithMock(t *testing.T) {
	cfg := &config.Database{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		Username: "test_user",
		Password: "test_password",
		Driver:   "postgres",
	}

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	mock.ExpectPing()

	_, err = Connect(cfg)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
