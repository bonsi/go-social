package main

import (
	"log"

	"github.com/bonsi/social/internal/db"
	"github.com/bonsi/social/internal/env"
	"github.com/bonsi/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:password@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewPostgresStorage(conn)
	db.Seed(store)
}
