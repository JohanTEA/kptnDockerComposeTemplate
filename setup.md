## KPTN - setup VPS from scratch

source: https://www.youtube.com/watch?v=F-9KWQByeU0

### Create VPS
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

### DNS name
- buy DNS name
- clear default ANC/DNS records
- lookup VPS IP address
ip addr
- point DNS name to vps IP adress
- wait for change to take effect (up to a day)
nslookup <dns.name>

### Setup VPS
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

### Optional
- install tmux (ssh session keep alive)
sudo apt install tmux
- logout
- login using tmux
tmux attach -t <kptn@ip-adress>
