package policy

import "database/sql"

type Policies interface {
	AddPolicy(*sql.Tx, *Policy) error
	Policies(*sql.Tx) ([]*Policy, error)
}

type Policy struct {
	ID string
}

type PolicyStorage struct {
}

func (st *PolicyStorage) AddPolicy(*sql.Tx, *Policy) error {
	return nil
}

func (st *PolicyStorage) Policies(*sql.Tx) ([]*Policy, error) {
	return nil, nil
}
