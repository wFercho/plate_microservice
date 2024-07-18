build:
	@go build -o bin/suppliers
run: build
	@./bin/suppliers
test:
	@go test -v ./...