#!/bin/bash
set -e

for app in kptn-main-frontend kptn-app1-frontend kptn-app1-backend; do
  cd ${app}
  ./dockerBuild.sh
  cd ..
done   

echo "Start Docker Compose according to 'docker-compose.yaml'"
sudo docker compose up -d
sudo docker compose ps

