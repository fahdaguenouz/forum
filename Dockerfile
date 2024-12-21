FROM golang:1.22 as builder

# Install necessary tools
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

# Set CGO_ENABLED=1
ENV CGO_ENABLED=1

# Set working directory
WORKDIR /app

# Copy files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o out

FROM debian:bullseye-slim

# Copy SSL certificates and the built application
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/out /app/

# Set working directory and command
WORKDIR /app
CMD ["./out"]
