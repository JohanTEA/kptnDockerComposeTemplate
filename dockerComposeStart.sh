#!/bin/bash
set -e

echo "Start Docker Compose according to 'docker-compose.yaml'"
sudo docker compose up -d
sudo docker compose ps

