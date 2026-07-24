package core_pgx_pool

import (
	"errors"

	core_postgres_pool "github.com/afkpanda2004/golang-todoapp/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (p pgxRow) Scan(dest ...any) error {
	err := p.Row.Scan(dest...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres_pool.ErrorNoRows
		}
		return err
	}

	return nil
}

type pgxCommantTag struct {
	pgconn.CommandTag
}
