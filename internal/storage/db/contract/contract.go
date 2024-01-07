package contract

import (
	"skeleton/internal/settings"
	"skeleton/internal/storage/db/policy"
	"skeleton/internal/storage/db/resource"
)

type Storage interface {
	policy.Policies
	resource.Resources
}

type Settings struct {
	Dialect settings.DBType
}
