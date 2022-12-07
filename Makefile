
build:
	go build -v ./...

test:
	go test -v ./...

lint:
	golangci-lint run -c ./.golangci.yml
