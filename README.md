Here's a description for the README file of your GitHub repository:

---

# Book Inventory API

This is a simple RESTful API for managing a book inventory system, built using Go and the Gin web framework. 
The API allows users to perform various operations such as retrieving the list of books, 
adding a new book, checking out a book, and returning a book.

## Features

-  Retrieve All Books : Fetch the list of all books in the inventory.
-  Retrieve a Book by ID : Get details of a specific book using its ID.
-  Add a New Book : Add a new book to the inventory.
-  Check Out a Book : Check out a book from the inventory.
-  Return a Book : Return a book to the inventory.

## Endpoints

### Get All Books
-  URL : `/books`
-  Method : `GET`
-  Description : Retrieves the list of all books in the inventory.
-  Curl Command : 
  ```sh
  curl localhost:8080/books
  ```

### Get Book by ID
-  URL : `/books/:id`
-  Method : `GET`
-  Description : Retrieves a book by its ID.
-  Curl Command : 
  ```sh
  curl localhost:8080/books/{id}
  ```

### Add a New Book
-  URL : `/books`
-  Method : `POST`
-  Description : Adds a new book to the inventory.
-  Curl Command : 
  ```sh
  curl -i -H "Content-Type: application/json" -d @body.json -X POST http://localhost:8080/books
  ```

### Check Out a Book
-  URL : `/checkout`
-  Method : `PATCH`
-  Description : Checks out a book from the inventory using its ID.
-  Curl Command : 
  ```sh
  curl -s -X PATCH 'localhost:8080/checkout?id={id}'
  ```

### Return a Book
-  URL : `/return`
-  Method : `PATCH`
-  Description : Returns a book to the inventory using its ID.
-  Curl Command : 
  ```sh
  curl -s -X PATCH 'localhost:8080/return?id={id}'
  ```

## Running the Application

To run the application, ensure you have Go installed on your machine. Clone the repository and execute the following commands:

```sh
go mod tidy
go run main.go
```

The server will start on `localhost:8080`.

## Example JSON Body for POST Request

To add a new book, create a `body.json` file with the following content:

```json
{
    "id": "4",
    "title": "1984",
    "author": "George Orwell",
    "quantity": 3
}
```

Use the following command to add the new book:

```sh
curl -i -H "Content-Type: application/json" -d @body.json -X POST http://localhost:8080/books
```

