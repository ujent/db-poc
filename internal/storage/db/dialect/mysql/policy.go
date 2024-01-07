package mysql

import (
	"database/sql"
	"fmt"
	"skeleton/internal/storage/db/policy"
)

// special realization for MySQL
// переопределяем метод
func (st *Storage) Policies(*sql.Tx) ([]*policy.Policy, error) {
	fmt.Println("MySQL Policies()")
	return nil, nil
}
