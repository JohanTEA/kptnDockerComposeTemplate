# KPTN - Docker Compose template with example Applications

Introduction to VPS running Docker Compose with Traefik routing, Nginx frontend and Go backend.

The routing in Traefik is set-up to comply with CORS rules for App1 frontend and App1 backend.

The guide assumes you run Linux (ex. Ubuntu) on your computer and have a basic understanding of Git.

- VPS OS - see [vpsOsSetup.md](vpsOsSetup.md)
- Docker Compose - see [docker-compose.md](docker-compose.md)
- Traefik for reverse proxy and load balancing
- Main FrontEnd - Nginx
- App1 FrontEnd - Nginx
- App1 BackEnd - Go

For application details:
- [docker-compose.yaml](docker-compose.yaml)

Developer notes for Nginx:
- [nginxInfo.md](nginxInfo.md)

## Goal

The goal of this template is to get you fully started locally. Setup of HTTPS certificates are not included in this template.

See instructions in linked guides on how to setup DNS name and HTTPS on your VPS with Traefik.

## Quick start

- Install Docker Compose - see [docker-compose.md](docker-compose.md)
- In terminal, run

    ./dockerBuildAll.sh
    ./dockerComposeStart.sh

- Go to http://localhost
- Change it - re-run it!

## TODO / Learn - Expand App1 with a database

At the moment there is no database in this system.

A suggestion to learn more is to expand App1 with a Postgres database [https://www.postgresql.org/](https://www.postgresql.org/).
- Expand Docker Compose with a Postgres database
- Expand App1 backend with database connectivity with API.s for App1 frontend.
- Expand App1 frontend with functionality to interact with App1 backend database functionality. 

## Release notes
- 2026jul22 - First release.
