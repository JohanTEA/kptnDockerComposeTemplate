#!/bin/bash
set -e

echo "Start Docker Compose according to 'docker-compose.yaml'"
sudo docker compose up -d --remove-orphans
sudo docker compose ps

