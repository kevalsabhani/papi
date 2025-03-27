# Pharmatail-Backend
- Backend server

## Project Structure

```
pharmatail-backend/
├── cmd/                      # Entry points for different services
│   └── api/                  # Main API server application
│       └── main.go
│
├── configs/                  # Configuration files (YAML, ENV files, etc.)
│   └── config.yaml
│
├── internal/                 # Application-specific packages (all business logic)
│   ├── auth/                 # Authentication & User Management (Includes RBAC)
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── rbac.go           # Role-Based Access Control logic
│   │
│   ├── subscription/         # Subscription Management (Plans, Payments, Renewals)
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   └── repository.go
│   │
│   ├── inventory/            # Raw Materials & Product Inventory
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   └── repository.go
│   │
│   ├── manufacturing/        # Manufacturing Process Management
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   └── repository.go
│   │
│   ├── orders/               # Order Management & Fulfillment
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   └── repository.go
│   │
│   ├── quality/              # Quality Control & Compliance
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   └── repository.go
│   │
│   ├── reporting/            # Reporting & Analytics
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── model.go
│   │   └── repository.go
│
├── pkg/                      # Common utilities and shared packages
│   ├── database/             # Database connection setup (PostgreSQL)
│   │   └── db.go
│   ├── logger/               # Centralized logging setup (Zap Logger)
│   │   └── logger.go
│   ├── middleware/           # Common middlewares (Authentication, CORS, etc.)
│   │   ├── auth_middleware.go
│   │   └── rbac_middleware.go
│   └── utils/                # Utility functions (JWT, Error handling, etc.)
│       ├── jwt.go
│       └── utils.go
│
├── migrations/               # Database migrations (SQL files)
│   ├── 001_init_schema.up.sql
│   └── 001_init_schema.down.sql
│
├── docs/                     # Documentation (API specs, architecture diagrams, etc.)
│
├── scripts/                  # Deployment and maintenance scripts (Docker, CI/CD, etc.)
│
├── go.mod                    # Go module file
├── go.sum                    # Dependency checksums
├── Makefile                  # Build, run, and clean commands
└── README.md                 # Project overview and documentation
```
