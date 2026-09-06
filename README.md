# common-libs

A shared Go library providing common infrastructure for bookmark management system.

## Overview

`common-libs` is designed to be reused across multiple services in a Clean Architecture.
It centralizes common concerns - authentication, rate limiting, database connectivity - 
so individual services can focus on business logic.

## Requirements

- Go 1.26+
- PostgreSQL
- Redis

## Installation

```bash
go get github.com/HemlockPham7/common-libs
```

## Key dependencies

| Purpose                     | Dependency                                             |
|-----------------------------|--------------------------------------------------------|
| Http framework              | github.com/gin-gonic/gin                               |
| JWT handling                | github.com/golang-jwt/jwt/v5                           |
| Rate limiting               | github.com/HemlockPham7/common-libs/pkg/ratelimitutils |
| New Relic                   | github.com/newrelic/go-agent/v3/newrelic               |
| Struct Validation           | github.com/go-playground/validator/v10                 |
| Envconfig                   | github.com/kelseyhightower/envconfig                   |
| CSV encoding/ decoding      | github.com/gocarina/gocsv                              |
| Structured logging          | github.com/rs/zerolog                                  |
| Redis Client                | github.com/redis/go-redis/v9                           |
| In-memory Redis for tests   | github.com/alicebob/miniredis/v2                       |
| Schema migrations           | github.com/golang-migrate/migrate/v4                   |
| PostgreSQL migration driver | github.com/golang-migrate/migrate/v4/database/postgres |
| File migration source       | github.com/golang-migrate/migrate/v4/source/file       |
| PostgreSQL driver           | gorm.io/driver/postgres                                |
| ORM                         | gorm.io/gorm                                           |
| SQLite driver for tests     | gorm.io/driver/sqlite                                  |
| GORM logger                 | gorm.io/gorm/logger                                    |
| Code hashing                | golang.org/x/crypto/bcrypt                             |
| Test assertion              | github.com/stretchr/testify/assert                     |

## Usage

Import the packages you need, for example:

```go
import (
	"github.com/HemlockPham7/common-libs/pkg/ratelimitutils"
	"github.com/HemlockPham7/common-libs/pkg/jwtutils"
	"github.com/HemlockPham7/common-libs/pkg/middlewares"
)
```

## Packages

| Package Name         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                               |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `pkg/array`          | Reusable utilities for working with slices/arrays, such as splitting collections into batches with minimal memory allocation.                                                                                                                                                                                                                                                                                                                             |
| `pkg/common`         | Provides reusable common utilities shared across different packages, such as simple error-handling helpers.                                                                                                                                                                                                                                                                                                                                               |
| `pkg/csv`            | Provides reusable CSV utilities for parsing CSV files, including files uploaded through multipart form requests.                                                                                                                                                                                                                                                                                                                                          |
| `pkg/dbutils`        | Provides reusable database utilities for normalizing and mapping database errors into application-level errors.                                                                                                                                                                                                                                                                                                                                           |
| `pkg/dto`            | Defines reusable Data Transfer Objects (DTOs) for structuring API request and response payloads, including standardized success responses and pagination metadata.                                                                                                                                                                                                                                                                                        |
| `pkg/errorutils`     | Provides reusable error handling utilities for normalizing and mapping errors into application-level errors.                                                                                                                                                                                                                                                                                                                                              |
| `pkg/jwtutils`       | Provides reusable utilities for generating and validating JSON Web Tokens (JWT) using RSA key pairs, including JWT signing with a private key and token validation with a public key.                                                                                                                                                                                                                                                                     |
| `pkg/logger`         | Provides reusable logging utilities for configuring the application's global log level.                                                                                                                                                                                                                                                                                                                                                                   |
| `pkg/middleware`     | Provides reusable Gin HTTP middleware for handling cross-cutting concerns such as JWT authentication and request rate limiting. <br /> `JWT authentication` — Extracts the Bearer token from the Authorization header, validates it, and stores the JWT claims in the Gin context. <br /> `Rate limiting` — Limits requests per user within a sliding time window and records rate-limit events in New Relic.                                             |
| `pkg/nrtrace`        | Provides reusable utilities for configuring and initializing a New Relic application using environment-based configuration.                                                                                                                                                                                                                                                                                                                               |
| `pkg/ratelimitutils` | Provides reusable Redis-based utilities for tracking and managing request rate limits, including incrementing and retrieving request counts.                                                                                                                                                                                                                                                                                                              |
| `pkg/redis`          | Provides reusable Redis utilities for configuring and creating Redis clients, including support for in-memory Redis clients for testing.                                                                                                                                                                                                                                                                                                                  |
| `pkg/requestutils`   | Provides reusable Gin request utilities for binding and validating request inputs, extracting authenticated user information from JWT claims, and handling request-related errors. <br /> `Request binding & validation` — Bind JSON body, URI parameters, query parameters, and headers into a struct, then validate the resulting input. <br /> `Authentication context` — Extract the authenticated user ID from JWT claims stored in the Gin context. |
| `pkg/response`       | Provides reusable API response structures and utilities for representing standardized application, input validation, authorization, and processing error responses.                                                                                                                                                                                                                                                                                       |
| `pkg/sqldb`          | Provides reusable SQL database utilities for configuring and creating GORM PostgreSQL clients, managing schema migrations, and initializing in-memory SQLite databases for testing.                                                                                                                                                                                                                                                                       |
| `pkg/utils`          | Provides reusable utility services for common application tasks, including secure random code generation and code hashing.                                                                                                                                                                                                                                                                                                                                |

## Development

### Run tests

```bash
make docker-test
```

Runs the tests in a Docker container, generates an HTML coverage report at `test-output/coverage.html`,
and exits with a non-zero status if the **coverage threshold** is not met, which is set to **90%**.