package main

import (
	"log"

	"github.com/bonsi/social/internal/db"
	"github.com/bonsi/social/internal/env"
	"github.com/bonsi/social/internal/store"
)

const version = "0.0.2"

//	@title			GopherSocial API
//	@description	API for GopherSocial, a social network for gophers
//	@termsOfService	http://swagger.io/terms

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8888"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8888"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_NAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
