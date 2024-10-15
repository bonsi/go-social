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
