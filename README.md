
# Kasir API

Kasir API is a simple RESTful API built with **Go** for a cashier (point-of-sale) system.  

This project is intended for learning and practice, demonstrating how to build a REST API using Go’s standard `net/http` package and how to document APIs using **Swagger** for easy testing and exploration.



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



## Swagger Documentation

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

You can use it to test all endpoints directly in your browser.
