# KPTN - App1 BackEnd

## Local development

Install Go

    sudo apt install go

Run Go app

    go run main.go

Publication on: http://127.0.0.1:8080/health

## Production

Build Go app for production

    mkdir build
    go build -ldflags "-s -w -X main.version=1.0" -o ./build/my-api

