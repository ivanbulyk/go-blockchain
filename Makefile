build:
	@go build -o bin/goblockchain

run: build
	@./bin/goblockchain

test:
	@go test -v ./...

.DEFAULT_GOAL := run