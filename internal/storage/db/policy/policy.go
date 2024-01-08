package policy

import (
	"context"
	"database/sql"
)

type Policies interface {
	AddPolicy(context.Context, *sql.Tx, *Policy) error
	Policies(context.Context, *sql.Tx) ([]*Policy, error)
}

type Policy struct {
	ID string
}

type PolicyStorage struct {
}

func (st *PolicyStorage) AddPolicy(context.Context, *sql.Tx, *Policy) error {
	return nil
}

func (st *PolicyStorage) Policies(context.Context, *sql.Tx) ([]*Policy, error) {
	return nil, nil
}
