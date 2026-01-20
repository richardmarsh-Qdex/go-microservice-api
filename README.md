# Go Microservice API

A microservice API built with Go, Gorilla Mux, and MongoDB.

## Features

- RESTful API endpoints
- JWT authentication
- User management
- Product management
- Order processing
- MongoDB integration

## Installation

```bash
go mod download
go run main.go
```

## API Endpoints

- GET /api/health - Health check
- POST /api/auth/register - Register user
- POST /api/auth/login - Login user
- GET /api/users - Get all users
- GET /api/products - Get all products
- POST /api/products - Create product
- GET /api/orders - Get all orders
- POST /api/orders - Create order
