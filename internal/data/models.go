package data

import (
	"errors"

	"github.com/edgedb/edgedb-go"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *edgedb.Client) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
