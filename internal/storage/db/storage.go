package storage

import (
	"fmt"
	"skeleton/internal/settings"
	"skeleton/internal/storage/db/contract"
	"skeleton/internal/storage/db/dialect/mysql"
	"skeleton/internal/storage/db/dialect/postgres"
	"skeleton/internal/storage/db/dialect/sqlite"

	_ "github.com/go-sql-driver/mysql"
)

func NewStorage(s *contract.Settings) (contract.Storage, error) {
	switch s.Dialect {
	case settings.MySQLDBType:
		return mysql.NewStorage(s)
	case settings.PostgresDBType:
		return postgres.NewStorage(s)
	case settings.SqliteDBType:
		return sqlite.NewStorage(s)
	default:
		return nil, fmt.Errorf("unexpected db type")
	}
}
