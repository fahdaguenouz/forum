# Stage 1: Build the Go application
FROM golang:1.22 as builder

# Install necessary tools
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

# Set CGO_ENABLED=1 for SQLite support
ENV CGO_ENABLED=1

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum, and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN go build -o out

# Stage 2: Create a lightweight runtime image
FROM debian:bullseye-slim

# Install CA certificates to enable HTTPS
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the SSL certificates, the built application, and the migration script
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/out /app/
COPY migrate.sh /app/migrate.sh

# Set the working directory inside the container
WORKDIR /app

# Make the migration script executable
RUN chmod +x /app/migrate.sh

# Set the entrypoint to run the migration script
ENTRYPOINT ["/app/migrate.sh"]

# Default command (this can be overridden at runtime)
CMD ["-h"]
