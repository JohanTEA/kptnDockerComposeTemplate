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

1. Install Docker Compose - see [docker-compose.md](docker-compose.md)
2. In terminal, run
```shell
./dockerBuildAll.sh
./dockerComposeStart.sh
```
3. Go to http://localhost
4. Change it - re-run it!

## TODO / Learn - Expand App1 with a database

At the moment there is no database in this system.

A suggestion to learn more is to expand App1 with a Postgres database [https://www.postgresql.org/](https://www.postgresql.org/).
- Expand Docker Compose with a Postgres database
- Expand App1 backend with database connectivity with API.s for App1 frontend.
- Expand App1 frontend with functionality to interact with App1 backend database functionality.

## TODO / Learn - Expand with automatic deployment

The system can be expanded with scheduled image checks and automatic deploy of updated apps. Good for safety and for easy deployment in production.

Recommended application for this is Watchtower [https://containrrr.dev/watchtower/](https://containrrr.dev/watchtower/).

## TODO / Learn - Expand with monitoring and alerts

As it is now the only monitoring in the system is the standard Docker Compose logging and the Traefik Dashboard. For a fully production ready system a proper monitoring and alert application is needed.

Here are some examples:
- Apprise: unified API to send notifications to Slack, Discord, Telegram, email and other services. It has no resourse or log monitoring and requires the applications to create alerts.
- Prometheus + Grafana + cAdvisor + Alertmanager: The industry standard for metrics and alerts. Prometheus scrapes metrics from targets, cAdvisor exposes container-level metrics, Grafana visualizes the data using pre-built dashboards, and Alertmanager handles alerts sent by client applications such as the Prometheus server and sends them to external integrations.


## Release notes

- 2026jul22 - First release.
