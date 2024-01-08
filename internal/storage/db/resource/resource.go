package resource

import (
	"context"
	"database/sql"
)

type Resources interface {
	AddResource(context.Context, *sql.Tx, *Resource) error
	Resourcies(context.Context, *sql.Tx) ([]*Resource, error)
}

type Resource struct {
	ID string
}

type ResourceStorage struct{}

func (st *ResourceStorage) AddResource(context.Context, *sql.Tx, *Resource) error {
	return nil
}

func (st *ResourceStorage) Resourcies(context.Context, *sql.Tx) ([]*Resource, error) {
	return nil, nil
}
