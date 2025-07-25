# Start from Go base image
FROM golang:latest

# Set working dir
WORKDIR /app

# Copy go.mod, go.sum จาก parent directory ของ Deploy
COPY ../go.mod ../go.sum ./

# Download dependencies
RUN go mod download

# Copy โค้ดที่เหลือจาก Deploy เข้าไป
COPY . .

# Build
RUN go build -o main .

# Expose port
EXPOSE 8000

# Run
CMD ["./main"]
