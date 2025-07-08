# To bootstrap common tasks

build:
	@go build -o bin/bank-api-go

run: build
	@./bin/bank-api-go

test:
	@go test -v ./...
