BINARY_NAME=gomux
GOPATH=$(shell go env GOPATH)
SHOULD_COMPRESS=$(COMPRESS)

tidy:
	go mod tidy

build:
	GOARCH=amd64 GOOS=linux go build -o bin/debug/${BINARY_NAME}-linux .
	GOARCH=amd64 GOOS=darwin go build -o bin/debug/${BINARY_NAME}-darwin .

build-prod:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/release/${BINARY_NAME}-linux .
	GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o bin/release/${BINARY_NAME}-darwin .
	if [ "${SHOULD_COMPRESS}" = "true" ]; then \
		upx -9 --best --lzma bin/release/${BINARY_NAME}-linux; \
		upx -9 --best --lzma bin/release/${BINARY_NAME}-darwin; \
	fi

install:
	if [ "${SHOULD_COMPRESS}" = "true" ]; then \
		go install -ldflags="-s -w" && upx -9 --best --lzma ${GOPATH}/bin/${BINARY_NAME}; \
	fi
	if [ "${SHOULD_COMPRESS}" != "true" ]; then \
		go install -ldflags="-s -w"; \
	fi

clean:
	go clean
	rm -f bin
