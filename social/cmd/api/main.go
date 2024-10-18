package main

import (
	"log"

	"github.com/bonsi/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8888"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
