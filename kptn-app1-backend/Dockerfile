# KPTN App1 BackEnd (GoLang)
# Images: https://hub.docker.com/

# Build: sudo docker build -t kptn-app1-backend .
# Run: sudo docker run -p 8091:8091 kptn-app1-backend

# Build stage
## golang:<golang version>-alpine
FROM golang:1.26-alpine AS builder
WORKDIR /appSrc
COPY . .
ENV CGO_ENABLED=0
RUN go build -C ./src -ldflags "-s -w -X main.version=$(cat version.txt)" -o /dist/kptnApp1BackEnd

# Runtime stage
## alpine:<alpine version>
FROM alpine:latest
COPY --from=builder /dist /app
# Optionally copy CA certs and timezone data
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app
USER appuser
ENTRYPOINT ["/app/kptnApp1BackEnd"]
