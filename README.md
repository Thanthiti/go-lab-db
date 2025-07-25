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

#Create  docs folder to add API document  
swag init 

#Authenthication
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v4

#Docker
# docker-compose
docker-compose up -d
docker-compose down -v

# Build Docker image
docker build -t my-app .

# Run container
docker run -p 8080:8080 --env-file .env my-app

```
## Note  

QueryRow use for return value
Exec use for no return