# Setup VPS from scratch

Reference: https://www.youtube.com/watch?v=F-9KWQByeU0 - YT, Dreams Of Code, Setting up a production ready VPS, Traefik

## Create VPS

ex. 'hostinger':
- base OS: ubuntu 24.04
- disable Monarx malware scanner
- create passw for root user
- add ssh key - todo...
- deploy VPS
- verify ssh connection
    ssh root@<ip-adress>
- add new user account
    adduser kptn
- add sudo privilege to new user (= add new user to 'sudo' group)
    usermod -aG sudo kptn

## DNS name

- buy DNS name
- clear default ANC/DNS records
- lookup VPS IP address
    ip addr
- point DNS name to vps IP adress
- wait for change to take effect (up to a day)
    nslookup <dns.name>

## Setup VPS

- login
    ssh kptn@<vps ip adress>
- set-up ssh key for kptn user
- disable root login, file: /etc/ssh/sshd_config
PermitRootLogin no
UsePAM no
- check other ssh conf files and repeat
- restart ssh service
    sudo systemctl reload ssh
- verify that root user is blocked
    ssh root@<vps ip-adress>

## Firewall

Rules:
- allow all outgoing requests by default
- deny all incoming requests by default
- allow incoming req. on port 22, open ssh
- allow incoming req. on port 80, http
- allow incoming req. on port 443, https

UFW - Uncomplicated Firewall, installed on Ubuntu by default.
If not installed, install:
    sudo apt install ufw

(!) Warning 1: Do all steps below before activating ufw, setting up ufw without allowing OpenSSH will kill your VPS.
(!) Warning 2: Docker Compose will override ufw/ip-tables if set-up incorrectly - reverse proxy needed with docker app config 'ports:' having '127.0.0.1:<port>' or Traefik labling instead.
    sudo ufw default allow outgoing
    sudo ufw default deny incoming
    sudo ufw allow OpenSSH
    sudo ufw allow 80
    sudo ufw allow 443
    
    sudo ufw show added
    sudo ufw status verbose
    sudo ufw enable

## Docker Compose

See dockerCompose for installation and running on VPS.

Docker Compose runs:
- Reverse Proxy (ex. Traefik) with TLS/HTTPS support
- Applications
- Other infrastructure tools as Watchtower, etc.

## Optional

- install tmux (ssh session keep alive)
    sudo apt install tmux
- logout
- login using tmux
    tmux attach -t <kptn@ip-adress>

