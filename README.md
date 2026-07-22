# KPTN - Docker Compose template with example Applications

Introduction to VPS running Docker Compose with Traefik routing, Nginx frontend and Go backend.

The guide assumes you run Linux (ex. Ubuntu) on the computer.

- VPS OS - see [vpsOsSetup.md]
- Docker Compose - see [docker-compose.md]
- Traefik for reverse proxy and load balancing
- Main FrontEnd - Nginx
- App1 FrontEnd - Nginx
- App1 BackEnd - Go

For application details:
- [docker-compose.yaml]

Developer notes for Nginx:
- [nginxInfo.md]

## Goal

The goal of this template is to get you fully started locally. Setup of HTTPS certificates are not included in this template.

See instructions in linked guides on how to setup DNS name and HTTPS on your VPS with Traefik.

## Quick start

- Install Docker Compose - see [docker-compose.md]
- In terminal, run 
    ./dockerBuildAll.sh
    ./dockerComposeStart.sh
- Go to http://localhost
- Change it - re-run it!

## Release notes
- 2026jul22 - First release.
