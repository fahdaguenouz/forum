# Stage 1: Build the Go application
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies for SSL and certificates
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && update-ca-certificates

# Copy Go source code into the container
COPY . .

# Install Go dependencies (if any)
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 go build -o out .

# Stage 2: Create the minimal runtime environment
FROM debian:bullseye-slim

# Install necessary libraries and dependencies for GLIBC
RUN apt-get update && apt-get install -y \
    libc6 \
    ca-certificates \
    && update-ca-certificates

# Copy SSL certificates and the built Go binary from the builder stage
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/out /app/

# Copy the migration script into the container
COPY migrate.sh /app/migrate.sh

# Set the working directory
WORKDIR /app

# Make the migrate script executable
RUN chmod +x /app/migrate.sh

# Set the environment variable for the Go application
ENV GO_ENV=production

# Expose the port the app will run on (use your app's port if different)
EXPOSE 8080

# Set the entrypoint to the migration and start script
CMD ["/app/migrate.sh"]
