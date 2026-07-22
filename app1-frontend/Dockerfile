# app1-frontend - App1 FrontEnd
## Reference: https://docs.docker.com/reference/dockerfile/
## Images: https://hub.docker.com/

## Build: sudo docker build -t app1-frontend .
## Run: sudo docker run -p 8081:80 app1-frontend

# Build stage
## nginx:<nginx version>-alpine
FROM nginx:1.31-alpine  AS builder
WORKDIR /appSrc
COPY . .
RUN cp -r ./src /dist
RUN mkdir /distNginx
RUN cp ./nginx/nginx.conf /distNginx/nginx.conf
RUN nginx -t -c /distNginx/nginx.conf  # Validate nginx.conf

# Runtime stage
## nginx:<nginx version>-alpine
FROM builder
COPY --from=builder /dist /usr/share/nginx/html
COPY --from=builder /distNginx/nginx.conf /etc/nginx/nginx.conf
# Optionally copy CA certs and timezone data
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# USER - The nginx image sets up user 'nginx' for runtime automatically
# ENTRYPOINT - The nginx image sets up entrypoint automatically

