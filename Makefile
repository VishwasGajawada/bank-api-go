# To bootstrap common tasks

build:
	@go build -buildvcs=false -o bin/bank-api-go

run: build
	@./bin/bank-api-go

test:
	@go test -v ./...
