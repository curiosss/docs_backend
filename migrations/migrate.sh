#!/bin/bash

# Load .env file
set -a
source .env
set +a

# Compose DB URL
DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Run migration
migrate -path migrations -database "$DB_URL" up
