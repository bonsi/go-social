package main

import (
	"log"
	"net/http"

	"github.com/bonsi/social/internal/store"
)

type commentKey string

const commentCtx commentKey = "comment"

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=100"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCommentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	comment := &store.Comment{
		PostID:  1,
		Content: payload.Content,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()
	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	log.Printf("%v\n", comment)
	log.Printf("%+v\n", comment)
	log.Printf("%#v\n", comment)

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
