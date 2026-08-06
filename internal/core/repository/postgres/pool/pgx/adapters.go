package core_pgx_pool

import (
	"errors"
	"fmt"

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

		return mapErrors(err)

	}

	return nil
}

type pgxCommantTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const (
		pgxViolatesForeignKey = "23503"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrorNoRows
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgxViolatesForeignKey {
			return fmt.Errorf(
				"%v, %w",
				err,
				core_postgres_pool.ErrorViolatesForeignKey,
			)

		}

	}
	return fmt.Errorf(
		"%v, %w",
		err,
		core_postgres_pool.ErrorUnknow,
	)
}
