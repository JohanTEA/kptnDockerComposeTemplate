#!/bin/bash
set -e

image_name=app1-backend

echo ""
echo "sudo docker run -p 8091:80 app1-backend"
sudo docker run -p 8091:80 app1-backend

