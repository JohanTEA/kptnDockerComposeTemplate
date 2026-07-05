# KPTN - App1 BackEnd

## Local development

Install Go
The snap argument '--classic' is so the application can access files outside the snap sandbox.

    sudo snap install go --classic

### Local build

Build Go app for production

    mkdir build
    go build -ldflags "-s -w -X main.version=$(cat version.txt).dev" -o ./build/kptnApp1BackEnd

### Run dev build

    go run ./build/kptnApp1BackEnd
    
Publication on: http://127.0.0.1:8091/health

## Production Docker build

    sudo docker build -t kptn-app1-backend .
    sudo docker images

### Run prod build

    sudo docker run -p 8091:8091 kptn-app1-backend
