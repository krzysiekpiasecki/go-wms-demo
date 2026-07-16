<h1 align="center">WMS API</h1>

<p align="center">
  <img src="assets/images/wms.pngalign="center">
  Warehouse Management System implemented in Go
</p>

---

Warehouse Management System (WMS) implemented in Go.

The project exposes a REST API for managing products and orders, using PostgreSQL as the persistence layer and Gin as the HTTP framework.

---

# Features

## Products

- Get product by ID

## Orders

- Create order
- Get order by ID
- Update order status
- Order items support

## Inventory

- Stock availability validation before order creation

## HTTP

- REST API based on Gin
- Request validation
- JSON responses
- Error handling

## Infrastructure

- PostgreSQL
- Repository pattern
- Service layer
- Middleware
- Structured logging with Zerolog
- OpenAPI specification

---

# Architecture

```text
HTTP Request
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
PostgreSQL
```

Project structure:

```text
cmd/
└── api/

internal/
├── config/
├── domain/
├── handler/
├── logger/
├── middleware/
├── repository/
│   └── postgres/
└── service/

migrations/
openapi/
assets/
└── images/
```

---

# Technologies

- Go
- Gin
- PostgreSQL
- PGX
- Zerolog
- OpenAPI 3.0

---

# Domain Model

## Product

```text
Product
├── ID
├── Name
└── CreatedAt
```

## Inventory

```text
Inventory
├── ProductID
└── Quantity
```

## Order

```text
Order
├── ID
├── Status
├── Comment
├── CreatedAt
└── Items
```

## Order Item

```text
OrderItem
├── ID
├── OrderID
├── ProductID
└── Quantity
```

---

# API Endpoints

## Health Check

### Request

```http
GET /health
```

### Response

```json
{
  "status": "ok"
}
```

---

## Get Product

### Request

```http
GET /products/{id}
```

### Response

```json
{
  "id": 1,
  "name": "Mug",
  "createdAt": "2026-07-15T11:54:56Z"
}
```

---

## Create Order

### Request

```http
POST /orders
```

```json
{
  "productId": 1,
  "quantity": 3,
  "comment": "urgent"
}
```

### Response

```http
201 Created
```

---

## Get Order

### Request

```http
GET /orders/{id}
```

### Response

```json
{
  "id": 3,
  "status": "NEW",
  "comment": null,
  "createdAt": "2026-07-16T14:46:28Z",
  "items": [
    {
      "id": 1,
      "orderId": 3,
      "productId": 1,
      "quantity": 3
    }
  ]
}
```

---

## Update Order Status

### Request

```http
PATCH /orders/{id}/status
```

```json
{
  "status": "COMPLETED"
}
```

### Response

```http
200 OK
```

---

# CURL Examples

## Health Check

```bash
curl http://localhost:8080/health
```

---

## Get Product

```bash
curl http://localhost:8080/products/1
```

---

## Create Order

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "productId": 1,
    "quantity": 3
  }'
```

---

## Create Order With Comment

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "productId": 1,
    "quantity": 3,
    "comment": "urgent"
  }'
```

---

## Get Order

```bash
curl http://localhost:8080/orders/3
```

---

## Update Order Status

```bash
curl -X PATCH http://localhost:8080/orders/3/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "COMPLETED"
  }'
```

---

## Product Not Found

```bash
curl http://localhost:8080/products/999
```

Response:

```json
{
  "error": "product not found"
}
```

---

## Order Not Found

```bash
curl http://localhost:8080/orders/999
```

Response:

```json
{
  "error": "order not found"
}
```

---

## Invalid Request

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{}'
```

---

## Insufficient Stock

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "productId": 1,
    "quantity": 99999
  }'
```

Response:

```json
{
  "error": "insufficient stock"
}
```

---

# Validation

Create order request validation:

```json
{
  "productId": 1,
  "quantity": 5
}
```

Rules:

- productId > 0
- quantity > 0

Invalid requests return:

```http
400 Bad Request
```

---

# Error Handling

## Product Not Found

```json
{
  "error": "product not found"
}
```

Response:

```http
404 Not Found
```

## Order Not Found

```json
{
  "error": "order not found"
}
```

Response:

```http
404 Not Found
```

## Insufficient Stock

```json
{
  "error": "insufficient stock"
}
```

Response:

```http
409 Conflict
```

---

# Middleware

## Logger

Logs every incoming request:

```text
INF http request method=GET path=/products/1 status=200
```

Example:

```text
2026-07-16T13:37:10+02:00 INF http request duration=0 method=POST path=/orders status=201
```

## Recovery

Recovers from panic and prevents application shutdown:

```json
{
  "error": "internal server error"
}
```

---

# Database

Main tables:

```sql
products
inventory
orders
order_items
```

Relations:

```text
orders
    │
    └── order_items
            │
            └── products
```

---

# Database Migration

Create database:

```sql
CREATE DATABASE wms;
```

Connect:

```bash
psql -U postgres -d wms
```

Apply migration scripts from the `migrations` directory in order:

```bash
psql -U postgres -d wms -f migrations/001_create_products.sql
psql -U postgres -d wms -f migrations/002_create_inventory.sql
psql -U postgres -d wms -f migrations/003_create_orders.sql
psql -U postgres -d wms -f migrations/004_create_order_items.sql
```

Verify tables:

```sql
\dt
```

Expected output:

```text
 inventory
 order_items
 orders
 products
```

Inspect schemas:

```sql
\d products
\d inventory
\d orders
\d order_items
```

---

# Running Locally

## Requirements

- Go 1.24+
- PostgreSQL

## Environment Variables

Create `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=wms
DB_SSLMODE=disable
```

## Run

```bash
go run ./cmd/api
```

Server:

```text
http://localhost:8080
```

---

# Tests

Run all tests:

```bash
go test ./...
```

Run service tests:

```bash
go test ./internal/service
```

Run handler tests:

```bash
go test ./internal/handler
```

---

# OpenAPI

OpenAPI specification:

```text
openapi/openapi.yaml
```

The specification describes:

- Products API
- Orders API
- Request schemas
- Response schemas
- Error responses

---

# Future Improvements

- Database transactions for order creation
- Inventory deduction after order creation
- Order status validation
- Swagger UI
- Docker Compose
- Connection pooling (pgxpool)
- JWT authentication
- Integration tests

---

WMS project created as a backend learning project focused on Go, REST API design, PostgreSQL, testing, and clean architecture.