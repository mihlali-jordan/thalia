package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
