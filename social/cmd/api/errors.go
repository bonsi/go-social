package main

import (
	"net/http"
)

func (app *application) forbiddenError(w http.ResponseWriter, r *http.Request) {
	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	writeJSONError(w, http.StatusForbidden, "forbidden")
}

func (app *application) unauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized basic error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem.")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("not found",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusNotFound, err.Error())
}

func (app *application) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf("conflict",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusConflict, err.Error())
}
