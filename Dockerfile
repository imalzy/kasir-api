FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN mkdir -p bin && \
    go build -o bin/kasir-api ./cmd/kasir-api

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/kasir-api .

# Expose port (adjust if needed)
EXPOSE 8080

# Run the application
CMD ["./kasir-api"]
