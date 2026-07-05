# NGINX info

**Note!** Do not use NGINX on your local computer.
Only use NGINX within a Docker container since it needs to take over the complete computer system.

## NGINX Conf
For the default NGINX configuration in a nginx:alpine Docker image, see 'nginxDefaultDockerImageConf.conf'.
- User 'nginx' must be used in Docker NGINX containers.

## NGINX commands
- verify nginx config, default path /etc/nginx/nginx.conf
    sudo nginx -t

- verify nginx config, non-default path
    sudo nginx -t -c /path/to/your/nginx.conf

- enable/start nginx server
config: /etc/nginx/nginx.conf
    sudo systemctl enable nginx
    sudo systemctl start nginx

- disable/stop nginx system start
    sudo systemctl disable nginx
    sudo systemctl stop nginx

- Check nginx service status
    sudo systemctl status nginx
    sudo systemctl is-enabled nginx
    ps -ef | grep -i nginx

