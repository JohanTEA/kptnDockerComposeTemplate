# KPTN Main FrontEnd
# Images: https://hub.docker.com/

# Build: sudo docker build -t kptn-main-frontend .
# Run: sudo docker run -p 8080:8080 kptn-main-frontend

# Build stage
## Using nginx image for building since the source is plain HTML.
## nginx:<nginx version>-apline
FROM nginx:1.31-alpine AS builder
WORKDIR /appSrc
COPY . .

# Runtime stage
FROM builder
COPY --from=builder /appSrc/src /usr/share/nginx/html
#COPY --from=builder /appSrc/nginx/nginx.conf /etc/nginx/nginx.conf
# Optionally copy CA certs and timezone data
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# USER - The nginx image sets up user 'nginx' for runtime
# ENTRYPOINT - The nginx image sets up entrypoint

