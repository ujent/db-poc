package mysql

import (
	"context"
	"database/sql"
	"skeleton/internal/storage/db/contract"
	"skeleton/internal/storage/db/policy"
	"skeleton/internal/storage/db/resource"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewStorage(s *contract.Settings) (contract.Storage, error) {
	db, err := sqlx.Connect("mysql", s.ConnStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db: db.DB}, nil
}

type Storage struct {
	once sync.Once
	db   *sql.DB

	*resource.ResourceStorage
	*policy.PolicyStorage
}

func (st *Storage) Close() error {
	var err error
	st.once.Do(func() {
		err = st.db.Close()
	})

	return err
}

func (st *Storage) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return st.db.BeginTx(ctx, opts)
}
