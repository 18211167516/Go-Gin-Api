# Go parameters
BINARY_NAME=api
BINARY_UNIX=$(BINARY_NAME)_unix
VERSION="v0.0.1"
DATE= `date +%Y%m%d%H%M%S`

.PHONY: all

all: version test build

version:
		@echo version: ${VERSION}

linux-build:

		@echo version: ${VERSION} date: ${DATE} os: linux
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_UNIX}

windows-build:
		@echo version: ${VERSION} date: ${DATE} os: windows
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}.exe	
test:
		go test tests/* -count=1
clean:
		go clean
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		go build -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

