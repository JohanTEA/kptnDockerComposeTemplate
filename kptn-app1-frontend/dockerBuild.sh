#!/bin/bash
set -e

image_name=kptn-app1-frontend
image_version=$(cat ./version.txt)

echo ""
echo "Building Docker image ${image_name}:${image_version}"
sudo docker build -t ${image_name}:${image_version} .
echo ""
echo "Updating Docker image ${image_name}:latest"
sudo docker tag ${image_name}:${image_version} ${image_name}:latest

