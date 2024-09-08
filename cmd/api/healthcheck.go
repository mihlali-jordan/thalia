package main

import (
	"fmt"
	"net/http"
)

// Implemented as a method on the application struct
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version %s\n", version)
}
