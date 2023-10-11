BINARY_NAME=./bin/order-pack

build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/main.go

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./... -c ./.golangci.yml