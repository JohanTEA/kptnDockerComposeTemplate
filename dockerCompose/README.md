# Docker Engine + Docker compose

https://hub.docker.com/
https://docs.docker.com/get-started/


## Install

https://docs.docker.com/engine/install/ubuntu/
https://docs.docker.com/compose/install/linux/

### Add Docker's official GPG key:

    sudo apt update
    sudo apt install ca-certificates curl
    sudo install -m 0755 -d /etc/apt/keyrings
    sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
    sudo chmod a+r /etc/apt/keyrings/docker.asc

### Add the repository to Apt sources:

    sudo tee /etc/apt/sources.list.d/docker.sources <<EOF
    Types: deb
    URIs: https://download.docker.com/linux/ubuntu
    Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
    Components: stable
    Architectures: $(dpkg --print-architecture)
    Signed-By: /etc/apt/keyrings/docker.asc
    EOF

### Install

    sudo apt update
    sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    sudo systemctl status docker

if docker is not running:
    sudo systemctl start docker

verify installation:
    sudo docker run hello-world

## Service commands

    sudo systemctl status docker

    sudo systemctl start docker
    sudo systemctl enable docker

    sudo systemctl stop docker
    sudo systemctl disable docker

## Test run

    sudo docker run hello-world

