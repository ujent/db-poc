package postgres

import (
	"skeleton/internal/storage/db/contract"
	"skeleton/internal/storage/db/policy"
	"skeleton/internal/storage/db/resource"
)

func NewStorage(Settings *contract.Settings) (contract.Storage, error) {
	return &Storage{}, nil
}

type Storage struct {
	*resource.ResourceStorage
	*policy.PolicyStorage
}
