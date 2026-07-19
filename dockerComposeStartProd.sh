#!/bin/bash
set -e

echo "Start Docker Compose according to 'docker-compose.prod.yaml'"
sudo docker compose -f docker-compose.prod.yaml up -d --remove-orphans
sudo docker compose ps

