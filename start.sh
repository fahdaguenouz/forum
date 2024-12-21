#!/bin/bash
echo "Running migrations..."
go run . -m
echo "Running seeders..."
go run . --seed
echo "Starting application..."
go run .
