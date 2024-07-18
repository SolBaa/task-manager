# Build stage
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Copy the .env file
COPY .env .env

# Build the Go app
RUN go build -o main ./cmd/server

# Run stage
FROM alpine:latest

WORKDIR /root/

# Install bash and netcat
RUN apk add --no-cache bash netcat-openbsd

# Copy the built binary and .env from the build stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Add the wait-for-it script
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Expose port (replace with the port your app uses)
EXPOSE 8080

# Command to run the binary with wait-for-it
CMD ["./wait-for-it.sh", "mysql:3306", "--", "./main"]
