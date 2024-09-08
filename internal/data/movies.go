package data

import (
	"context"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/mihlali-jordan/thalia/internal/validator"
)

type Movie struct {
	ID        edgedb.UUID `edgedb:"id"`
	CreatedAt time.Time   `edgedb:"created_at"`
	Title     string      `edgedb:"title"`
	Year      int32       `edgedb:"year"`
	Runtime   Runtime     `edgedb:"runtime"`
	Genres    []string    `edgedb:"genres"`
	Version   int32       `edgedb:"version"`
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

	args := []interface{}{movie.Title, movie.Year, movie.Runtime, movie.Genres, movie.Version}
	return m.DB.QuerySingle(context.Background(), query, &inserted, args)
}

func (m MovieModel) Get(id edgedb.UUID) error {
	return nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id edgedb.UUID) error {
	return nil
}
