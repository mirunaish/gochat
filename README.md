# Go Chatroom

## What you built?

I built a chatroom application. Users can create an account and join the
chatroom and communicate in real time with other users. Messages are not saved,
so they disappear after a few seconds, making the experience closer to talking
in-person.

### Server

The focus of this project was the Go backend - a websocket server that enables
bidirectional communication. It connects to a Postgres database that holds
accounts.

### Client

A minimal client designed to connect to the server. Open multiple
instances of it to set up a connection.

Include some screenshots.
[How?](https://help.github.com/articles/about-readmes/#relative-links-and-image-paths-in-readme-files)

## Who Did What?

I did everything.

## What you learned

I learned:
- the module / package structure of Go projects and how to import files
- how to create REST routes using Gin
- error checking, how to create custom errors, handling different error types
- Go structs and interfaces and how they can be used to mimic OOP
- GORM, a GO ORM

What worked:
- 

What didn't work:
- I initially tried gorilla/mux, but that was more difficult to use so I switched to Gin.

## Authors

Miruna Palaghean

## Acknowledgments

### Documentation
- https://go.dev/ref/spec
- https://go.dev/doc/effective_go
- https://go.dev/doc/code
- https://gorm.io/docs/index.html

### Tutorials

- https://go.dev/doc/tutorial/getting-started
- https://gowebexamples.com/routes-using-gorilla-mux/
- https://go.dev/doc/tutorial/web-service-gin
- https://gowebexamples.com/password-hashing/
- https://dev.to/karanpratapsingh connecting-to-postgresql-using-gorm-24fj

### Other

- various online sources, linked in the code where they were used