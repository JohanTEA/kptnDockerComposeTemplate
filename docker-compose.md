# Docker and Docker Compose 
This includes Docker Engine etc.

https://hub.docker.com/
https://docs.docker.com/get-started/

Docker Compose intro
- https://github.com/docker/getting-started-todo-app/tree/main

- https://www.youtube.com/watch?v=eGz9DS-aIeY (YT, NetworkChuck, Learn Docker)
- https://www.youtube.com/watch?v=DM65_JyGxCo (YT, NetworkChuck, Learn Docker Compose)
- https://www.youtube.com/watch?v=F-9KWQByeU0 (YT, Dreams Of Code, Setting up a production ready VPS, Traefik)
-- https://github.com/dreamsofcode-io/guestbook
-- https://doc.traefik.io/traefik/getting-started/
-- https://containrrr.dev/watchtower/
- https://www.youtube.com/watch?v=CmUzMi5QLzI (Jim's Garage, Traefik 3.3)
-- https://github.com/JamesTurland/JimsGarage/tree/main/Traefikv3
- https://www.youtube.com/watch?v=mZbLvGQqEvY (Andrew Schmelyun, Manage multiple sites with Docker and Traefik)
-- https://github.com/aschmelyun/docker-compose-traefik

## Installation (Local development and VPS server)

https://docs.docker.com/engine/install/ubuntu/
https://docs.docker.com/compose/install/linux/

### Add Docker's official GPG key:

    sudo apt update
    sudo apt install ca-certificates curl
    sudo install -m 0755 -d /etc/apt/keyrings
    sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
    sudo chmod a+r /etc/apt/keyrings/docker.asc

### Add the repository to Apt sources

```shell
    sudo tee /etc/apt/sources.list.d/docker.sources <<EOF
    Types: deb
    URIs: https://download.docker.com/linux/ubuntu
    Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
    Components: stable
    Architectures: $(dpkg --print-architecture)
    Signed-By: /etc/apt/keyrings/docker.asc
    EOF
```

### Install

    sudo apt update
    sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    sudo systemctl status docker

if docker is not running:
    sudo systemctl start docker

verify installation:
    sudo docker run hello-world

## systemctl commands

    sudo systemctl status docker

    sudo systemctl start docker
    sudo systemctl enable docker

    sudo systemctl stop docker
    sudo systemctl disable docker

## Docker Test run

### Docker test image 'hello-world'

    sudo docker run hello-world

### Base nginx image

    sudo docker system info
    sudo docker pull nginx
    sudo docker run --name nginx-test -itd -p 8081:80 nginx
    sudo docker ps
    sudo docker rm nginx-test

## Docker commands

    docker --help

### Managing local Docker image repo

    sudo docker images
    sudo docker pull <image from Docker hub>
    sudo docker rmi <image id>
    
### Managing local Docker containers

    sudo docker container ls
    sudo docker container stop <container id>
    sudo docker container rm <container id>

    sudo docker container prune #removes all stopped containers

    sudo docker exec -it <container_name_or_id> /bin/sh # shell access for Alpine images
    sudo docker exec -it <container_name_or_id> /bin/bash # shell access for Ubuntu images

## Docker Compose

    sudo docker compose up -d  # Starts DC 'docker-compose.yaml'
    
    sudo docker-compose logs -f  # Shows standard-out logs for all apps

    sudo docker compose ps  # Lists DC containers

    sudo docker compose stop  # Stops DC 'docker-compose.yaml' containers

    sudo docker compose down  # Stops and prunes DC 'docker-compose.yaml' containers
    sudo docker compose down --remove-orphans  # Stops and prunes DC all containers

## Docker Network

    sudo docker network ls  # List networks
    
    sudo docker network prune  # Remove All Unused Networks
    
