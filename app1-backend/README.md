# App1 BackEnd

## Local development

Install Go.

The snap argument '--classic' is so the application can access files outside the snap sandbox.

    sudo snap install go --classic

### Local build

Build Go app for production

    mkdir build
    go build -C ./src -ldflags "-s -w -X main.version=$(cat version.txt).dev" -o ./build/app1BackEnd

### Run dev build

    go run ./build/app1BackEnd
    
Publication on: http://127.0.0.1:8091/health
