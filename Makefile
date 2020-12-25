# Go parameters
BINARY_NAME=api
VERSION="v0.0.1"
DATE= `date +%Y%m%d%H%M%S`
CMD="gga"

.PHONY: all cmd

all: version test build

cmd: 
		@echo "make gga 构建gga终端工具"
		go build -o ${CMD} cmd/*

version:
		@echo version: ${VERSION}

linux-build:
		@echo version: ${VERSION} date: ${DATE} os: linux
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}

windows-build:
		@echo version: ${VERSION} date: ${DATE} os: windows
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}.exe	
test:
		go test tests/* -count=1
clean:
		go clean
		@if [ -f ${CMD} ] ; then rm ${CMD} ; fi
		@if [ -f ${BINARY_NAME} ] ; then rm ${BINARY_NAME} ; fi
		@if [ -f ${BINARY_NAME}.exe ] ; then rm ${BINARY_NAME}.exe ; fi
run:
		go build -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

