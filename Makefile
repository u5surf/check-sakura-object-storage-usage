.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: lint
lint:
	golangci-lint run -v ./...
	go vet ./...
