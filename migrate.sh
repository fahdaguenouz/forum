#!/bin/bash

# Run database migrations
echo "Running database migrations..."
# Run the migration command using the built executable
./out migrate

# Check if migration was successful
if [ $? -eq 0 ]; then
  echo "Database migrations completed successfully."
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
