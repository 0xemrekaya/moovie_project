# Stage 1: Build the application
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./cmd/api

# Stage 2: Create the final, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the migrations. The application needs them to run on startup.
COPY --from=builder /app/migrations ./migrations

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["/app/server"] 