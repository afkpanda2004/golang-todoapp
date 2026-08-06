package core_postgres_pool

import "errors"

var (
	ErrorNoRows             = errors.New("no rows")
	ErrorViolatesForeignKey = errors.New("violates foreign key")
	ErrorUnknow             = errors.New("unknow")
)
