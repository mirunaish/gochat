# Go Chatroom

## What you built?

I had two goals for this assignment: 
1. learn a new programming language. I chose Go because it has been increasing
in popularity and it seemed somewhat different from other programming languages
I know.
2. deploy a dockerized application on AWS. Previously my only deployment
experience was in Render, and AWS and Docker seem widely used.

I built a simple chatroom application. Users can create an account and join the
chatroom and communicate in real time with other users. Messages are not saved,
so they disappear after a few seconds, making the experience closer to talking
in-person.

### Server

The focus of this project was the Go backend - a websocket server that enables
bidirectional communication. It connects to a Postgres database that holds
accounts.

![generics](media/first_ping.png)
*first ping*

### Client

A minimal client designed to connect to the server. Open multiple
instances of it to set up a connection.

Include some screenshots.
[How?](https://help.github.com/articles/about-readmes/#relative-links-and-image-paths-in-readme-files)

## Running instructions

1. Create a postgres database. Tables will be created automatically when first running the program.
2. Generate a secret key for signing JWTs.
3. Duplicate `.env.template` and rename it to `.env`. Fill in the values.
4. `cd server`
5. Run `go run ./cmd/`

## Who Did What?

I did everything. (see Acknowledgments section for links to tutorials)

## What you learned

I learned:
- the module / package structure of Go projects and how to import files
- how to create REST routes and middleware using Gin
- generic functions and type parameters in Go
- error checking in Go, how to create custom errors, handling different error types
- Go structs and interfaces and how they can be used to mimic OOP
- GORM, a GO ORM
- how to color console output

![generics](media/generics.png)
*image source: [Type parameters in Go](https://bitfieldconsulting.com/posts/type-parameters)*

What worked:
- everything, eventually, hopefully

What didn't work:
- I initially tried gorilla/mux, but that was more difficult to use so I switched to Gin.
- I had difficulty setting up the formatter provided by the VS Code Go extension.
- I made my AWS account more than a year ago so my free tier expired.
- I tried signing up for Github Student for potential free DataOcean, but they still haven't processed my application.
- I signed up for Mogenius because they promised free deployment but it was a lie.
- I made an Oracle Cloud account because they promised a free database but they thought I gave them false information (?).

## Authors

Miruna Palaghean

## Acknowledgments

### Documentation
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [How to Write Go Code](https://go.dev/doc/code)
- [GORM](https://gorm.io/docs/index.html)
- [golang-jwt](https://golang-jwt.github.io/jwt/usage/create/)

### Tutorials

- [Get started with Go](https://go.dev/doc/tutorial/getting-started)
- [Routing (using gorilla/mux)](https://gowebexamples.com/routes-using-gorilla-mux/)
- [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
- [Password Hashing](https://gowebexamples.com/password-hashing/)
- [Connecting to PostgreSQL using GORM](https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj)
- [Gin Custom Middleware](https://gin-gonic.com/docs/examples/custom-middleware/)

### Other

- various online sources that i glanced at briefly, linked in the code where they were used