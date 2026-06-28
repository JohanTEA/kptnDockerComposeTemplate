#!/bin/bash
set -e

NGINX_CONF=$(pwd)/fe/nginx/nginx.local.conf
echo Validating nginx conf...
sudo nginx -t -c $NGINX_CONF
sudo systemctl stop nginx
sudo systemctl start nginx
sudo nginx -c $NGINX_CONF
echo ""
echo Nginx service started http://127.0.0.1:80/ http://127.0.0.1:8080/
echo ""
ps -ef | grep -i nginx

