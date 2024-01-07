package contract

import (
	"context"
	"database/sql"
	"skeleton/internal/settings"
	"skeleton/internal/storage/db/policy"
	"skeleton/internal/storage/db/resource"
)

type Storage interface {
	policy.Policies
	resource.Resources

	Close() error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Settings struct {
	Dialect settings.DBType
	ConnStr string
}
