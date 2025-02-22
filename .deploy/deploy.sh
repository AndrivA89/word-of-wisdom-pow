#!/bin/bash
# .deploy/deploy.sh
#
# Deployment script for Word of Wisdom application.
# This script builds Docker images, stops old containers,
# and starts the application using Docker Compose.
#
# Requirements:
# - Docker Desktop (with Docker Compose v2)
# - Project files in the current working directory
#
# Usage: ./deploy.sh

set -e
set -o pipefail

echo "Starting deployment process..."

# Build Docker images
echo "Building Docker images..."
docker compose build

# Stop and remove existing containers
echo "Stopping and removing old containers..."
docker compose down

# Start containers in detached mode
echo "Starting containers..."
docker compose up -d

echo "Deployment completed successfully!"
