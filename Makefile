# Go parameters
BINARY_NAME=go-api
BINARY_UNIX=$(BINARY_NAME)_unix
VERSION="v0.0.1"
DATE= `date +%Y%m%d%H%M%S`

.PHONY: all

all: version test build

version:
		@echo version: ${VERSION}
build:
		@echo version: ${VERSION} date: ${DATE} os: windows
		@go build -o $(BINARY_NAME) -v -race
test:
		go test tests/* -count=1
clean:
		go clean
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		go build -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

# Cross compilation
build-linux:
		@echo version: ${VERSION} date: ${DATE} os: linux-centOS
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) -v