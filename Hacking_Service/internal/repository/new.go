package repository

import (
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	Matrix
}

func New(db *pg.DB) *Repository {
	return &Repository{NewRepository(db)}
}
