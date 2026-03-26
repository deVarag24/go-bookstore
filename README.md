# Go Bookstore API

This is a simple RESTful API for managing a bookstore, built using the Go programming language and the Fiber web framework. The API allows you to perform CRUD (Create, Read, Update, Delete) operations on books in the bookstore.


## Features
- Create a new book
- Retrieve a list of all books
- Retrieve a single book by ID
- Update an existing book
- Delete a book by ID

## Technologies Used
- Go (Golang)
- Fiber web framework
- PostgreSQL (for database)
- GORM (for ORM)
- Docker (for containerization)

## API Endpoints
- `POST /api/v1/books`: Create a new book
- `GET /api/v1/books`: Retrieve a list of all books
- `GET /api/v1/books/{id}`: Retrieve a single book by ID
- `PUT /api/v1/books/{id}`: Update an existing book
- `DELETE /api/v1/books/{id}`: Delete a book by ID


