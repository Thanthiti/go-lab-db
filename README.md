# go-lab-db

Practice repo for PostgreSQL and GORM in Go.

## Packages used

- `database/sql`
- `github.com/lib/pq` — PostgreSQL driver
- `github.com/joho/godotenv` — Load .env files
- `github.com/gofiber/swagger` swagger
- `gorm.io/gorm/logger` - logger 

## Setup

```bash
go mod init myModule
go get github.com/lib/pq
go get github.com/joho/godotenv 
go get github.com/gofiber/fiber/v2

# ORM Package 
go get -u gorm.io/gorm 
go get -u gorm.io/driver/postgres
# Swagger for APIs
go get github.com/swaggo/swag
go get github.com/gofiber/swagger
swag init
```
## Note 
QueryRow use for return value
Exec use for no return