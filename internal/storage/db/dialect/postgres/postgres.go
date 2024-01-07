package postgres

import (
	"context"
	"database/sql"
	"skeleton/internal/storage/db/contract"
	"skeleton/internal/storage/db/policy"
	"skeleton/internal/storage/db/resource"
)

func NewStorage(s *contract.Settings) (contract.Storage, error) {
	return &Storage{}, nil
}

type Storage struct {
	db *sql.DB

	*resource.ResourceStorage
	*policy.PolicyStorage
}

func (st *Storage) Close() error {
	return st.db.Close()
}

func (st *Storage) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return st.db.BeginTx(ctx, opts)
}
