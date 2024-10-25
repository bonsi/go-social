# Backend Engineering with Go

## Section 1: Introduction

### 1. Project Overview

### 2. Why Go with Go?

### 3. Preface for Udemy Students

### 4. Course Resources

- <https://github.com/sikozonpc/GopherSocial>

### 5. Getting your Tools Ready

Starting on Module 3. (Scaffolding our API Server) we'll start using Go
specific concepts such as: context, interfaces, error handling, pointers,
Goroutines, Channels & Maps.

In this course, I won’t be diving into the basics of Go (that’s coming in the
future). However, I’ve created separate videos on these topics, and I’d
recommend going through them to learn the fundamental concepts before moving
forward.

You can also refer back to these resources and explore them as you come across
relevant topics during the course. That’s what I tend to do—I prefer learning
things as I need them.

Resources:

- Context: <https://youtu.be/Q0BdETrs1Ok>
- Error Handling: <https://youtu.be/dKUiCF3abHc>
- Interfaces: <https://youtu.be/4OVJ-ir9hL8?si=nZcSoQrTXrYh69y4>
- Maps: <https://youtu.be/999h-iyp4Hw?si=fPLtWRs7DWIVBIk->
- Pointers: <https://youtu.be/DVNOP1LE3Mg?si=KXaKeHeIipjLg1HZ>
- Goroutines & Channels: <https://youtu.be/3QESpVGiiB8?si=kqpETtKp73Abyiyw>

## Section 2: Project Architecture

### 6. Design Principles for a REST API

- <https://12factor.net/>
- "Roy Fielding - REST dissertation", see `./resources/fielding_dissertation.pdf`
- <https://martinfowler.com/articles/richardsonMaturityModel.html>

## Section 3: Building a server from TCP to HTTP

### 7. TCP Server - net package

- <https://pkg.go.dev/net>

### 8. Understanding Routing

### 9. HTTP Server - The net/http package

### 10. Encoding & Decoding JSON Requests

## Section 4: Scaffolding our API Server

### 11. Setting up your Development Environment

- Complete backend API in Golang (JWT, MySQL & Tests): <https://www.youtube.com/watch?v=7VLmLOiQ3ck>

```sh
mkdir social
cd social

go mod init github.com/bonsi/social

mkdir -p {bin,cmd/migrate/migrations,cmd/api,docs,internal,scripts,web}
```

### 12. Clean Layered Architecture

- <https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164>

### 13. Setting up the HTTP server and API

```sh
cd social
go get -u github.com/go-chi/chi/v5
```

### 14. Hot Reloading in Go

- <https://github.com/air-verse/air>

### 15. Environment Variables

- Go-specific package: <https://github.com/joho/godotenv>
- Generic CLI tool: <https://direnv.net/>

- <https://12factor.net/config>

```sh
# direnv is automatically loaded from ~/.zshrc

cd social

# allow direnv to read our .envrc
# NOTE: we need to do this again every time we change .envrc
direnv allow .
```

## Section 5: Databases

### 16. The Repository Pattern

- <https://threedots.tech/post/repository-pattern-in-go/>

### 17. Implementing the Repository Pattern

### 18. Persisting data with SQL

Some ORM suggestions if we do not want to use the Go standard lib below (lib/pq):

- <https://github.com/jmoiron/sqlx>
- <https://github.com/volatiletech/sqlboiler>

```sh
go get github.com/lib/pq
```

### 19. Configuring the DB Connection Pool

```sh
docker compose up
```

### 20. SQL Migrations

- <https://github.com/golang-migrate/migrate/tree/master/cmd/migrate>
- <https://www.postgresql.org/docs/current/citext.html>
- <https://www.gnu.org/software/make/manual/make.html>

```sh
cd social

# create empty migrations files (up & down)
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users

# migrate
migrate -path=./cmd/migrate/migrations -database="postgres://postgres:password@localhost/socialnetwork?sslmode=disable" up
```

These commands in a Makefile:

```sh
make migration posts_create

make migration add_users_fk_to_posts_table

make migrate-up
```

## Section 6: Posts CRUD

### 21. Marshalling JSON responses

```sh
curl http://localhost:8888/v1/health -v
```

### 22. Creating a User Feed Post

```sh
make migration alter_posts_with_tags_updated

make migrate-up
```

- POST request to `http://localhost:8888/v1/posts` using Insomnia fails
- FK error, manually insert user using SQL for now
- creating a post now works

### 23. Getting a Post From a User

### 24. Internal Errors Package
