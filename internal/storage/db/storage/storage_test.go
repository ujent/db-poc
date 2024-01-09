package storage

import (
	"context"
	"skeleton/internal/settings"
	"skeleton/internal/storage/db/contract"
	"skeleton/internal/storage/db/policy"
	"skeleton/internal/storage/db/resource"
	"testing"
)

// usage example
func TestAddPolicy(t *testing.T) {
	st, err := NewStorage(&contract.Settings{Dialect: settings.MySQLDBType})
	if err != nil {
		t.Fatal(err)
	}

	defer st.Close()

	ctx := context.Background()

	tx, err := st.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	err = st.AddPolicy(ctx, tx, &policy.Policy{})
	if err != nil {
		t.Fatal(err)
	}

	err = st.AddResource(ctx, tx, &resource.Resource{})
	if err != nil {
		t.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
