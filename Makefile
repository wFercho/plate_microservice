build:
	@go build -o bin/plate_microservice
run: build
	@./bin/plate_microservice
test:
	@go test -v ./...