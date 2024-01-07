package mysql

import (
	"fmt"
	"skeleton/internal/storage/db/policy"
)

// special realization for MySQL
// переопределяем метод
func (st *Storage) Policies() ([]*policy.Policy, error) {
	fmt.Println("MySQL Policies()")
	return nil, nil
}
