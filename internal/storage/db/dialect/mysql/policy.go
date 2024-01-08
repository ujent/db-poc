package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"skeleton/internal/storage/db/policy"
)

// special realization for MySQL
// переопределяем метод
func (st *Storage) Policies(context.Context, *sql.Tx) ([]*policy.Policy, error) {
	fmt.Println("MySQL Policies()")
	return nil, nil
}
