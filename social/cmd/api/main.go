package main

import (
	"log"

	"github.com/bonsi/social/internal/env"
	"github.com/bonsi/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8888"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
