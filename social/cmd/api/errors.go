package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error (method: %s, path: %s): %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem.")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error (method: %s, path: %s): %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error (method: %s, path: %s): %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, err.Error())
}

func (app *application) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Conflict error (method: %s, path: %s): %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusConflict, err.Error())
}
