# Backend Engineering with Go

- <https://www.udemy.com/course/backend-engineering-with-go/>
- <https://gitlab.com/bonsi/backend-engineering-with-go>

## Project info

- API with Postgres database

- Postgres 16 (postgres/password)
- Docker Compose
- Air (local, not via docker)
- mprocs to run services during development
- direnv
- Insomnia

- <http://localhost:8888>
- Health check:

  - <http://localhost:8888/v1/health>

  ```sh
  curl -v http://localhost:8888/v1/health
  ```

- Posts:

  - Create:

    - <http://localhost:8888/v1/posts>

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

### 25. HTTP Payload Validation

- <https://github.com/go-playground/validator>

```sh
go get github.com/go-playground/validator/v10
```

### 26. DB Relationships & SQL Joins

- <https://www.dataquest.io/blog/sql-joins/>
- SQL Cheatsheet: ./resources/DB+Relationships+SQL+Joins+Backend+Engineering+with+Go.jpg

  ![SQL Cheatsheet](./resources/DB+Relationships+SQL+Joins+Backend+Engineering+with+Go.jpg)

### 27. Adding Comments to Posts

```sh
cd social
make migration add_comments
```

### 28. Updating and Deleting Posts

### 29. Standardising JSON Responses

### 30. Optimistic Concurrency Control

- <https://cult.honeypot.io/reads/optimistic-vs-pessimistic-concurrency/>

```sh
make migration add_version_to_posts
# ...
make migrate-up
```

- also added `app.conflictError`, which returns HTTP 409, to inform the user
  of a version conflict while trying to update a post
- see `./social/scripts/test_concurrency.go` for the code used to test this
  error

### 31. Managing SQL Query Timeouts

### 32. Database Seeding

```sh
make seed

2025/01/06 00:03:33 pq: SSL is not enabled on the server
exit status 1
make: *** [Makefile:19: seed] Error 1
```

- the instructor advised to remove the double quotes around DB_ADDR in `.envrc`

- the instructor made some booboos in `seed.go` by picking values from the wrong
  slices. Fixed that.
- also added a new Makefile command `make migrate-drop` to drop the whole database

### 33. Database Seeding: "SSL not enabled" problem

Some students are getting the "SSL not enabled" problem.

Here are some common solutions to that problem:

1. Remove the surrounding quotes from your .envrc variables.
2. Make sure add the ?sslmode=disable at the end of the connection string.
3. Docker image of postgres is conflicting with an older one -> Delete volumes
   and containers.

### Assignment 1: Challenge: Create Comments for Posts

## Section 7: User Feed

### 34. Creating the User Profile

### 35. Adding Followers

```sh
make migration add_followers_table

# ...

make migrate-up
```

### 36. SQL Indexes

- <https://niallburkley.com/blog/index-columns-for-like-in-postgres/>
- <https://www.postgresql.org/docs/13/indexes-intro.html>
- <https://www.postgresql.org/docs/9.5/pgtrgm.html>

```sql
-- to get the read time
EXPLAIN ANALYSE SELECT * FROM comments WHERE content ILIKE '%post%';

-- the comment search query
SELECT * FROM comments WHERE content like ILIKE '%post%';

-- inserting many rows
do $$
begin
for r in 1..100000 loop
INSERT INTO comment (post_id, user_id, content) VALUES (749, 400, 'Super duper OMG ... So cool COOL!');
end loop;
end;
$$;

-- the index to increase read speed at the expense of write speed (faster reads, slower writes)
CREATE INDEX idx_comments_content ON comments USING gin (content gin_trgm_ops);
```

```sh
make migration add_indexes

# ...

make migrate-up
```

### 37. User Feed Algorithm

---

---

IB: Somwhere during following this course, a new section was added:
Section 3: Mini Course: Advanced Go
(Already did those lectures in the other course: "Complete Microservices with Go")

Therefore, the section and lecture numbers jump ahead from this point on

---

## Section 9: Filtering, Sorting, and Pagination

### 51. Pagination and Sorting

### 52. Feed Filtering

## Section 10: Documentation

### 53. Windows problems when generating swagger docs

### 54. Auto Generating Docs for the API

- <https://github.com/swaggo/swag?tab=readme-ov-file>
- <https://github.com/swaggo/http-swagger>

```sh
go install github.com/swaggo/swag/cmd/swag@latest

swag init

go get -u github.com/swaggo/http-swagger
```

```sh
make gen-docs
```

- <http://localhost:8888/v1/swagger/index.html>
- docs will be auto-generated through the Air pre-cmd

### 55. Documenting the Handlers

## Section 11: Structured logging

### 56. Adding a Logger

- <https://pkg.go.dev/go.uber.org/zap#section-readme>

```sh
cd ./social/

go get -u go.uber.org/zap
```

## Section 12: User creation

### 57. User Registration Overview

```sh
cd ./social/

make migration add_invitation

make migration add_activated_to_user
```

```sh
make migrate-up
```

### 58. SQL Transactions

```sh
make migration add_expiry_to_invitations

# ...

make migrate-up
```

### 59. User Activation

## Section 13: Sending emails

### 60. Sending the Invitation Email

- <https://gobyexample.com/embed-directive>
- <https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go>

- I created a SendGrid account (owned by Twilio now apparantly) but got locked
  out or whatever. Don't care, switching to MailTrap in the next lecture

- after adding new env vars:

```sh
direnv allow .
```

### 61. (Optional) Sending Emails using MailTrap

### 62. Improving Further the Email

### 63. Extra: Building the Confirmation UI

```sh
cd ./web/

npm create vite@latest . -- --template react-ts

npm i

npm install react-router-dom

```

- `cd ./web/ && npm run dev`
- <http://localhost:5173/>
- <http://localhost:5173/confirm/TOKEN>

```sh
go get github.com/go-chi/cors
```

## Section 14: Authentication

### 64. Authentication Overview

### 65. Basic Authentication

### 66. Generating Tokens

```sh
curl -X 'POST' \
  'http://localhost:8888/v1/authentication/token' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "test2@example.com",
  "password": "secret"
}'

> {
>  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnb3BoZXJzb2NpYWwiLCJleHAiOjE3NTU1MDEyMDksImlhdCI6MTc1NTI0MjAwOSwiaXNzIjoiZ29waGVyc29jaWFsIiwibmJmIjoxNzU1MjQyMDA5LCJzdWIiOjEwMn0.5fF6j-KRWlgcGwa2qhd7JUuHqSPFnfO8HabvtVof2dM"
> }
```

- we can see the contents of the token by pasting it here: <https://www.jwt.io/>

### 67. Validating Tokens

## Section 15: Authorization

### 68. Authorization Overview

### 69. Authorization Database Setup

```sh
make migration add_roles_table

make migration alter_users_with_roles
```

```sh
make migrate-up
```

- the seeding did not work anymore because of the added role_id to the users
  table. So I added it to the model, the seed method, et voila.

### 70. Role Precedence Middleware

### 71. Fixing the User Invitation
