# Stage 1: Build the Go application
FROM golang:1.23 as builder

WORKDIR /app

# Install necessary dependencies for SSL and certificates
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && update-ca-certificates

# Copy Go source code into the container
COPY . .

# Ensure Go dependencies are installed
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 go build -o out .

# Stage 2: Create the minimal runtime environment
FROM debian:bullseye-slim

# Install necessary libraries
RUN apt-get update && apt-get install -y \
    libc6 \
    ca-certificates \
    && update-ca-certificates

# Copy SSL certificates and the built Go binary from the builder stage
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/out /app/

# Copy the Front-end/views directory
COPY --from=builder /app/Front-end/views /app/Front-end/views

# Copy the migration script
COPY migrate.sh /app/migrate.sh

# Set the working directory
WORKDIR /app

# Make the migration script executable
RUN chmod +x /app/migrate.sh

# Set the environment variable
ENV GO_ENV=production

# Expose the application port
EXPOSE 8080

CMD ["/app/migrate.sh"]
