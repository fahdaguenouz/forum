#!/bin/bash

# Function to check the migration status
run_migration() {
  # Run the migration based on the Go application's output
  echo "Running migrations..."
  ./out
}

# Check if database migration is required
echo "Running database migrations..."

# Run the Go application to handle migrations
run_migration

# Check if migration was successful
if [ $? -eq 0 ]; then
  echo "Database migrations completed successfully!"
else
  echo "Database migrations failed!"
  exit 1
fi

# Run database seeding if '--seed' flag is passed
if [[ "$@" == *"--seed"* ]]; then
  echo "Seeding the database..."
  # Run the seeding command using the built executable
  ./out seed
fi

# Start the Go application (if migration and seeding are successful)
echo "Starting the application..."
exec ./out "$@"
