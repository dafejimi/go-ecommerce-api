# E-commerce API

A RESTful API for an e-commerce platform built with Go, featuring user authentication, product management, and order processing.

## Features

- JWT Authentication and RBAC
- Product Management (CRUD)
- Order Processing
- MySQL Database with GORM
- RESTful API Design
- Input Validation

## Prerequisites

- Go 1.21+
- MySQL 8.0+

## Project Structure
```
ecommerce-api/
├── cmd/
│   └── main.go                  # Application entry point
├── config/
│   └── config.go                # Configuration setup
├── controllers/
│   ├── user_controller.go
│   ├── product_controller.go
│   └── order_controller.go
├── middleware/
│   └── auth_middleware.go       # JWT authentication
├── models/
│   ├── user.go
│   ├── product.go
│   └── order.go
├── repositories/
│   ├── user_repository.go
│   ├── product_repository.go
│   └── order_repository.go
├── services/
│   ├── auth_service.go
│   ├── user_service.go
│   ├── product_service.go
│   └── order_service.go
├── utils/
│   ├── hash.go
│   └── jwt.go
└── docs/
    └── swagger.yaml
```

## API Endpoints

### Auth
- `POST /api/register` - Register user
- `POST /api/login` - Login user

### Products
- `GET /api/products` - List products
- `GET /api/products/:id` - Get product
- `POST /api/products` - Create product (Admin)
- `PUT /api/products/:id` - Update product (Admin)
- `DELETE /api/products/:id` - Delete product (Admin)

### Orders
- `POST /api/orders` - Create order
- `GET /api/orders` - List user orders
- `GET /api/orders/:id` - Get order details
- `PUT /api/orders/:id/status` - Update status (Admin)
- `PUT /api/orders/:id/cancel` - Cancel order

## Setup

1. Clone and install dependencies:
```bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
go mod download
```

2. Configure environment variables:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=ecommerce
JWT_SECRET=your_secret_key
PORT=8080
```

3. Run:
```bash
go run cmd/main.go
```
