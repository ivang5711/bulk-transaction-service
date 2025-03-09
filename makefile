build:
	@go build -o bin/bts

run: build
	@./bin/bts

test:
	@go test -v ./...