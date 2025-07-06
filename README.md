# go-lab-db

Practice repo for PostgreSQL and GORM in Go.

## Packages used

- `database/sql`
- `github.com/lib/pq` — PostgreSQL driver
- `github.com/joho/godotenv` — Load .env files

## Setup

```bash
go mod init myModule
go get github.com/lib/pq
go get github.com/joho/godotenv 
go get github.com/gofiber/fiber/v2 

```
## Note f
QueryRow use for return value
Exec use for no return