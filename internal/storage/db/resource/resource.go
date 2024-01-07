package resource

import "database/sql"

type Resources interface {
	AddResource(*sql.Tx, *Resource) error
	Resourcies(*sql.Tx) ([]*Resource, error)
}

type Resource struct {
	ID string
}

type ResourceStorage struct{}

func (st *ResourceStorage) AddResource(*sql.Tx, *Resource) error {
	return nil
}

func (st *ResourceStorage) Resourcies(*sql.Tx) ([]*Resource, error) {
	return nil, nil
}
