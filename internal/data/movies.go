package data

import (
	"context"
	"errors"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/mihlali-jordan/thalia/internal/validator"
)

var edbErr edgedb.Error

type Movie struct {
	ID        edgedb.UUID          `edgedb:"id" json:"id"`
	CreatedAt edgedb.LocalDateTime `edgedb:"created_at" json:"created_at"`
	Title     string               `edgedb:"title" json:"title"`
	Year      int32                `edgedb:"year" json:"year"`
	Runtime   int32                `edgedb:"runtime" json:"runtime"`
	Genres    []string             `edgedb:"genres" json:"genres"`
	Version   int32                `edgedb:"version" json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least one genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

type MovieModel struct {
	DB *edgedb.Client
}

func (m MovieModel) Insert(movie *Movie) error {
	var inserted struct{ id edgedb.UUID }
	query := `
		INSERT Movie {
			title := <str>$0,
			year := <int32>$1,
			runtime := <int32>$2,
			genres := <array<str>>$3
		}
	`

	args := []interface{}{movie.Title, movie.Year, movie.Runtime, movie.Genres}
	return m.DB.QuerySingle(context.Background(), query, &inserted, args...)
}

func (m MovieModel) Get(id edgedb.UUID) (*Movie, error) {
	var movie Movie
	query := `
		SELECT Movie {
			id,
			title,
			year,
			runtime,
			genres,
			version
		} filter .id = <uuid>$0
	`

	err := m.DB.QuerySingle(context.Background(), query, &movie, id)
	if err != nil {
		switch {
		case errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &movie, nil
}

func (m MovieModel) Update(movie *Movie) error {
	var updated struct{ id edgedb.UUID }
	query := `
		update Movie
		filter .id = <uuid>$0
		set { title := <str>$1, year := <int32>$2, runtime := <int32>$3, genres := <array<str>>$4}
	`
	args := []interface{}{movie.ID, movie.Title, movie.Year, movie.Runtime, movie.Genres}
	return m.DB.QuerySingle(context.Background(), query, &updated, args...)
}

func (m MovieModel) Delete(id edgedb.UUID) error {
	var deleted struct{ id edgedb.UUID }
	query := `
		delete Movie
		filter .id = <uuid>$0
	`
	return m.DB.QuerySingle(context.Background(), query, &deleted, id)
}
