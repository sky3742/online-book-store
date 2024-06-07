# Online Book Store Backend

This is a backend server for an online bookstore built with Golang. It provides functionalities for user authentication, listing books, placing orders, and retrieving order histories.

## Features

- **User Authentication**: Users can create an account using email and password.
- **Book Listing**: Retrieve a list of books available in the store.
- **Order Management**: Users can place orders for books and retrieve their order histories.
- **Session Management**: Authentication is managed via session tokens. A valid session token is required to place orders.

## Technology Stack

- **Language**: Golang
- **Web Framework**: [Fiber](https://github.com/gofiber/fiber)
- **Database**: SQLite
- **Session Management**: Custom session token-based authentication

## Prerequisites

- Go (1.18 or later)
- SQLite

## Getting Started

1. **Clone the repository**:
    ```sh
    git clone https://github.com/sky3742/online-book-store.git
    cd online-book-store
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Setup the database**:
    Use the provided Makefile to manage the database.

    - **Migrate**:
        ```sh
        make migrate
        ```

    - **Rollback**:
        ```sh
        make rollback
        ```

    - **Seed**:
        ```sh
        make seeds
        ```

4. **Run the server**:
    ```sh
    make start
    ```

5. **Watch for changes**:
    ```sh
    make watch
    ```

6. **Run tests**:
    ```sh
    make test
    ```

## Makefile Commands

- **migrate**: Run database migrations
- **rollback**: Rollback the last database migration
- **seeds**: Seed the database with initial data
- **start**: Start the server
- **watch**: Use Air to serve the server. Any changes made will rebuild the server without restarting it.
- **mock**: Generate mock interfaces using mockery
- **test**: Run tests

## API Endpoints

### User Authentication

- **Register**: `POST /api/v1/register`
  - Request: 
  ```json
  { 
    "email": "user@example.com", 
    "password": "password123"
  }
  ```
  - Response: 
  ```json
  {
    "status_code": 201,
    "message": "success register account",
    "data": {
      "token": "da10bd65-b0b7-4736-a6f1-21f75a9b17e1"
    }
  }
  ```

- **Login**: `POST /api/v1/login`
  - Request:
  ```json
  { 
    "email": "user@example.com",
    "password": "password123"
  }
  ```
  - Response: 
  ```json
  {
    "status_code": 200,
    "message": "success login account",
    "data": {
      "token": "dafabcca-422f-4f87-b494-5d83f6611e02"
    }
  }
  ```

### Books

- **Get Books**: `GET /api/v1/books`
  - Response:
  ```json
  {
    "status_code": 200,
    "message": "success get books",
    "data": [
      {
        "id": "3e95296c-6289-4d02-a187-a514c8d55c77",
        "created_at": "2024-06-07T18:45:04.543547+08:00",
        "updated_at": "2024-06-07T18:45:04.543547+08:00",
        "title": "To Kill a Mockingbird",
        "description": "A classic novel depicting racial injustice in the American South.",
        "author": "Harper Lee",
        "publication_year": 1960,
        "genre": [
          "Fiction",
          "Classic"
        ],
        "cover_image": "https://fakeimg.pl/667x1000/cc6600",
        "price": 123.45
      },
    ]
  }
  ```

- **Get Book**: `GET /api/v1/books/<book_id>`
  - Response: 
  ```json
  {
    "status_code": 200,
    "message": "success get book",
    "data": {
      "id": "3e95296c-6289-4d02-a187-a514c8d55c77",
      "created_at": "2024-06-07T18:45:04.543547+08:00",
      "updated_at": "2024-06-07T18:45:04.543547+08:00",
      "title": "To Kill a Mockingbird",
      "description": "A classic novel depicting racial injustice in the American South.",
      "author": "Harper Lee",
      "publication_year": 1960,
      "genre": [
        "Fiction",
        "Classic"
      ],
      "cover_image": "https://fakeimg.pl/667x1000/cc6600",
      "price": 123.45
    }
  }
  ```

### Orders

- **Place Order**: `POST /api/v1/orders`
  - Request: `X-Session-Token: <session_token>`
  - Body:
  ```json
  {
    "items": [
      {
        "book_id": "3e95296c-6289-4d02-a187-a514c8d55c77",
        "quantity": 4
      }
    ]
  }
  ```
  - Response:
  ```json
  {
    "status_code": 201,
    "message": "success create order",
    "data": {
      "id": "432be7f9-4f1d-4696-89ca-a255195e9da2",
      "created_at": "2024-06-08T10:08:03.517954+08:00",
      "updated_at": "2024-06-08T10:08:03.517954+08:00",
      "items": [
        {
          "id": "fd8e96d9-4205-4f0a-9d0b-84df424869b0",
          "created_at": "2024-06-08T10:08:03.519503+08:00",
          "updated_at": "2024-06-08T10:08:03.519503+08:00",
          "book_id": "3e95296c-6289-4d02-a187-a514c8d55c77",
          "price": 123.45,
          "quantity": 4
        }
      ],
      "status": "pending",
      "total_price": 493.8
    }
  }
  ```

- **Get Order History**: `GET /orders`
  - Request: `X-Session-Token: <session_token>`
  - Response:
  ```json
  {
    "status_code": 200,
    "message": "success get orders",
    "data": [
      {
        "id": "432be7f9-4f1d-4696-89ca-a255195e9da2",
        "created_at": "2024-06-08T10:08:03.517954+08:00",
        "updated_at": "2024-06-08T10:08:03.517954+08:00",
        "items": [
          {
            "id": "fd8e96d9-4205-4f0a-9d0b-84df424869b0",
            "created_at": "2024-06-08T10:08:03.519503+08:00",
            "updated_at": "2024-06-08T10:08:03.519503+08:00",
            "book_id": "3e95296c-6289-4d02-a187-a514c8d55c77",
            "price": 123.45,
            "quantity": 4
          }
        ],
        "status": "pending",
        "total_price": 493.8
      }
    ]
  }
  ```

- **Get Order Detail**: `GET /orders/<order_id>`
  - Request: `X-Session-Token: <session_token>`
  - Response:
  ```json
  {
    "status_code": 200,
    "message": "success get orders",
    "data": {
      "id": "432be7f9-4f1d-4696-89ca-a255195e9da2",
      "created_at": "2024-06-08T10:08:03.517954+08:00",
      "updated_at": "2024-06-08T10:08:03.517954+08:00",
      "items": [
        {
          "id": "fd8e96d9-4205-4f0a-9d0b-84df424869b0",
          "created_at": "2024-06-08T10:08:03.519503+08:00",
          "updated_at": "2024-06-08T10:08:03.519503+08:00",
          "book_id": "3e95296c-6289-4d02-a187-a514c8d55c77",
          "price": 123.45,
          "quantity": 4
        }
      ],
      "status": "pending",
      "total_price": 493.8
    }
  }
  ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
