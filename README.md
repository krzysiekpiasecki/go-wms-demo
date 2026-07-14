# go-wms-demo
 Simple Warehouse Management System (WMS) demo implemented as a Go monolith with a database backend. 
 
 The application supports inbound stock operations, order creation and validation based on available inventory, outbound shipments, and inventory reporting. Includes unit tests with code coverage above 50%.

## Project Structure

Initial project layout following Clean Architecture principles.

```text
.
├── cmd/
│   └── api/                    # Application entrypoint
├── internal/
│   ├── config/                 # Configuration loading and management
│   ├── domain/                 # Domain models and business rules
│   ├── handler/                # HTTP handlers/controllers
│   ├── repository/
│   │   └── postgres/           # PostgreSQL repository implementations
│   └── service/                # Application/business services
├── migrations/                 # Database migration files
└── test/                       # Integration and end-to-end tests
