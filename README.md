
# Produk API

Produk API is a simple RESTful API for managing products, built with Go. It allows you to create, read, update, and delete products. The API also includes Swagger documentation for easy testing and exploration.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)
- [Response Format](#response-format)
- [Swagger Documentation](#swagger-documentation)
- [Health Check](#health-check)

## Features

- Get all products
- Get a product by ID
- Add a new product
- Update an existing product by ID
- Delete a product by ID
- JSON-formatted responses with consistent structure
- Swagger UI documentation

## Tech Stack

- **Language:** Go
- **Router:** net/http
- **Swagger:** [swaggo/http-swagger](https://github.com/swaggo/http-swagger)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/produk-api.git
cd produk-api
````

2. Install dependencies:

```bash
go mod tidy
```

3. Generate Swagger docs (if needed):

```bash
swag init
```

> Make sure you have [Swag CLI](https://github.com/swaggo/swag) installed.

## Running the API

Start the server:

```bash
go run main.go
```

The API will run on `http://localhost:8080`.

## API Endpoints

### Produk

| Method | Endpoint           | Description            |
| ------ | ------------------ | ---------------------- |
| GET    | `/api/produk`      | Get all products       |
| POST   | `/api/produk`      | Add a new product      |
| GET    | `/api/produk/{id}` | Get a product by ID    |
| PUT    | `/api/produk/{id}` | Update a product by ID |
| DELETE | `/api/produk/{id}` | Delete a product by ID |

### Health Check

| Method | Endpoint  | Description             |
| ------ | --------- | ----------------------- |
| GET    | `/health` | Check if API is running |

## Request & Response Format

### Example Product JSON

```json
{
  "id": 1,
  "nama": "Indomie Godog",
  "harga": 3500,
  "stok": 10
}
```

### Standard API Response

```json
{
  "status": 200,
  "data": {},
  "message": "Success message"
}
```

## Swagger Documentation

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

You can use it to test all endpoints directly in your browser.
