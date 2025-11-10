# Build stage
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with FIPS mode enabled
# The GODEBUG=fips140=on environment variable activates FIPS 140-3 mode in Go 1.24+
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd/server

# Runtime stage
FROM debian:bookworm-slim

# Install ca-certificates for HTTPS connections
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Set environment variables for FIPS mode
ENV GODEBUG=fips140=on

# Expose port
EXPOSE 4000

# Run the server
CMD ["./server"]
