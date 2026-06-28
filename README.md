## KPTN

Set-up VPS setup.md


### Local development tools
- tmux
sudo apt install tmux

- Go compiler
sudo snap install go --classic

- nginx (web + reverse proxy server)
sudo apt install nginx


### FE - nginx
- verify nginx config
sudo nginx -t -c /path/to/your/nginx.conf

- start nginx server for testing (nginx service must already be running)
sudo nginx -c /path/to/your/nginx.conf

- enable/start nginx server for production as a service (Docker image)
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

### BE - Go

