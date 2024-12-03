# ===== Stage 1: Build the Go Binary =====
FROM golang:1.23-alpine AS builder

# Install Git (required for some Go modules)
RUN apk update && apk add --no-cache git

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
# -ldflags="-s -w" strips the binary for smaller size
RUN go build -o app -ldflags="-s -w" main.go

# ===== Stage 2: Create the Final Image =====
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the port your application listens on
EXPOSE 8085

# Command to run the executable
CMD ["./app serve"]