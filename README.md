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

Import the packages you need:

```go
import (
	"github.com/HemlockPham7/common-libs/pkg/ratelimitutils"
	"github.com/HemlockPham7/common-libs/pkg/jwtutils"
	"github.com/HemlockPham7/common-libs/pkg/middlewares"
)
```

## Packages

### `pkg/array`

Reusable utilities for working with slices/arrays, such as splitting collections into batches with minimal memory allocation.

### `pkg/common`

### `pkg/csv`

### `pkg/dbutils`

### `pkg/dto`

### `pkg/errorutils`

### `pkg/jwtutils`

### `pkg/logger`

### `pkg/middleware`

### `pkg/nrtrace`

### `pkg/ratelimitutils`

### `pkg/redis`

### `pkg/requestutils`

### `pkg/response`

### `pkg/sqldb`

### `pkg/utils`

## Development

### Run tests

```bash
make docker-test
```

Runs the tests in a Docker container, generates an HTML coverage report at `test-output/coverage.html`,
and exits with a non-zero status if the **coverage threshold** is not met, which is set to **90%**.