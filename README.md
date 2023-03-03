# GO Boilerplate
[![Go Reference](https://pkg.go.dev/badge/github.com/andhikayuana/qiscus-unofficial-go.svg)](https://pkg.go.dev/github.com/syahidfrd/go-boilerplate)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/andhikayuana/qiscus-unofficial-go)](https://goreportcard.com/report/github.com/syahidfrd/go-boilerplate)

### Directory Structure
```
├── cmd
│   └── api             // Main applications for this project
├── config              // Configuration app
├── delivery            // Delivery usefull to encode raw body, query params, make a response to client
│   ├── http
│   └── middleware
├── docs                // Design and user documents (in addition to your godoc generated documentation)
├── domain              // Domains is a struct to store any objects from database, example you have a products in your tables.
├── infrastructure
│   └── datastore       // Database configuration
├── migration           // Database migrations.
├── mocks               // Mock code, generated by mockery
├── repository          // Represents the behavior of the data store. It can be implemented using PostgreSQL, MongoDB, etc
│   ├── pgsql
│   └── redis
├── transport           // Request and response model
│   └── request
├── usecase             // Usecase contains all of business logic
└── utils               // Additional services
    ├── crypto
    ├── jwt
    └── logger
```

### Prerequisite and full list what has been used
* [PostgreSQL](https://www.postgresql.org) - PostgreSQL Database
* [Redis](https://redis.com) - Redis
* [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library
* [echo](https://github.com/labstack/echo) - Web framework
* [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang
* [zap](https://github.com/uber-go/zap) - Logger
* [ozzo-validation](github.com/go-ozzo/ozzo-validation) - Go Struct and Field validation
* [mockery](https://github.com/vektra/mockery) - A mock code autogenerator for Golang
* [swag](https://github.com/swaggo/swag) - Generator RESTful API documentation with Swagger
* [echo-swagger](https://github.com/swaggo/echo-swagger) - Echo middleware for Swagger
* [testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks
* [godotenv](github.com/joho/godotenv) - Environment variables loader

App requires 2 database (PostgreSQL and Redis server), run from your local machine or run it using docker with the following command
```
# run postgreSQL
docker run -d -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go_boilerplate postgres

# run redis
docker run -d -p 6379:6379 redis
``` 

### Migration
Run below command to run migration
```
migrate -path migration -database "${DATABASE_URL}" up
```

To create a new migration file
```
migrate create -ext sql -dir migration -seq name
```

### Test
Run below command to run test, and make sure that all tests are passing
```
go test -v ./...
```

### Running
Run below command to run app
```
go run ./cmd/api/main.go
```

Swagger URL
```
${BASE_URL}/swagger/index.html
```

You can find usefull commands in `Makefile`.
