package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) link.Repository {
	return repository{
		db: db,
	}
}
